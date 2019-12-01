package beater

import (
	"fmt"
	"io/ioutil"
	s "strings"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/seungyongshim/lsbeat/config"
)

// Lsbeat configuration.
type Lsbeat struct {
	done          chan struct{}
	config        config.Config
	client        beat.Client
	period        time.Duration
	paths         []string
	smbDrives     []string
	lastIndexTime time.Time // 가장 마지막 검색한 시간
}

// New creates an instance of lsbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Lsbeat{
		done:   make(chan struct{}),
		config: c,
	}

	bt.paths = c.Paths
	bt.smbDrives = c.SmbDrives

	return bt, nil
}

// Run starts lsbeat.
func (bt *Lsbeat) Run(b *beat.Beat) error {
	logp.Info("lsbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	d := bt.config.Period
	ticker := time.NewTicker(d)
	for {
		for _, path := range bt.paths {
			path = s.Replace(path, "\\", "/", -1)
			listDir(path, bt, b)
		}

		for _, path := range bt.smbDrives {
			path = s.Replace(path, "\\", "/", -1)
			checkSmbDrive(path, bt, b)
		}

		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
	}
}

// Stop stops lsbeat.
func (bt *Lsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func checkSmbDrive(dirName string, bt *Lsbeat, b *beat.Beat) {
	_, err := ioutil.ReadDir(dirName)
	if err != nil {
		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"dirName": dirName,
				"isExist": false,
				"smb":     true,
				"reason":  "\\\\192.168.100.76\\watch",
			},
		}
		bt.client.Publish(event)
	}
}

func listDir(dirName string, bt *Lsbeat, b *beat.Beat) (int64, int, int) {
	files, err := ioutil.ReadDir(dirName)

	if err != nil {
		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"dirName": dirName,
				"isExist": false,
			},
		}
		bt.client.Publish(event)
		return 0, 0, 0
	}

	fileinfos := []common.MapStr{}

	dirSize := int64(0)
	dirSizeAcc := int64(0)
	filecount := 0
	filecountAcc := 0
	subfoldercount := 0
	subfoldercountAcc := 0

	for _, f := range files {
		if f.IsDir() {
			subfoldercount++
			da, fa, sfa := listDir(dirName+"/"+f.Name(), bt, b)
			dirSizeAcc += da
			filecountAcc += fa
			subfoldercountAcc += sfa
		} else {
			dirSize += f.Size()

			fileinfos = append(fileinfos, common.MapStr{
				"fileName": f.Name(),
				"fileSize": f.Size(),
				"modTime":  f.ModTime(),
			})
			filecount++
		}
	}

	dirSizeAcc += dirSize
	filecountAcc += filecount
	subfoldercountAcc += subfoldercount

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"files":                    fileinfos,
			"filesCount":               filecount,
			"filesCountAccumulate":     filecountAcc,
			"dirName":                  dirName,
			"dirSize":                  dirSize,
			"dirSizeAccumulate":        dirSizeAcc,
			"subFolderCount":           subfoldercount,
			"subFolderCountAccumulate": subfoldercountAcc,
			"isExist":                  true,
		},
	}
	bt.client.Publish(event)
	logp.Info("lsbeat published listDir.")
	return dirSizeAcc, filecountAcc, subfoldercountAcc
}
