package beater

import (
	"fmt"
	"io/ioutil"
	"os"
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
	depth         int
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
	bt.depth = c.Depth

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
			if path == "nothing" {
				continue
			}
			path = s.Replace(path, "\\", "/", -1)
			listDir(path, bt, b, bt.depth)
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

func listDir(dirName string, bt *Lsbeat, b *beat.Beat, depth int) (int64, int, int) {
	curDir, err := os.Stat(dirName)

	if err != nil {
		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"dirName":  dirName,
				"isExist":  false,
				"isFolder": true,
			},
		}
		bt.client.Publish(event)
		return 0, 0, 0
	}

	if depth == 0 {
		return curDir.Size(), 0, 0
	}

	files, _ := ioutil.ReadDir(dirName)

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
			da, fa, sfa := listDir(dirName+"/"+f.Name(), bt, b, depth-1)
			dirSizeAcc += da
			filecountAcc += fa
			subfoldercountAcc += sfa
		} else {
			dirSize += f.Size()

			fileinfos = append(fileinfos, common.MapStr{
				"name":    f.Name(),
				"size":    f.Size(),
				"modTime": f.ModTime(),
			})
			filecount++
		}
	}

	dirSizeAcc += dirSize
	filecountAcc += filecount
	subfoldercountAcc += subfoldercount

	folderinfo := common.MapStr{
		"name":     curDir.Name(),
		"size":     dirSize,
		"modTime":  curDir.ModTime(),
		"depth":    bt.depth - depth,
		"maxDepth": bt.depth,
	}

	eventDir := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"folder":                   folderinfo,
			"dirName":                  dirName,
			"dirSize":                  dirSize,
			"dirSizeAccumulate":        dirSizeAcc,
			"filesCount":               filecount,
			"filesCountAccumulate":     filecountAcc,
			"subFolderCount":           subfoldercount,
			"subFolderCountAccumulate": subfoldercountAcc,
			"isExist":                  true,
			"isFolder":                 true,
		},
	}

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"files":    fileinfos,
			"dirName":  dirName,
			"isExist":  true,
			"isFolder": false,
		},
	}
	bt.client.Publish(event)
	bt.client.Publish(eventDir)
	return dirSizeAcc, filecountAcc, subfoldercountAcc
}
