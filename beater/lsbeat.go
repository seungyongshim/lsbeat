package beater

import (
	"fmt"
	"io/ioutil"
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

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		for _, path := range bt.paths {
			listDir(path, bt, b)
		}
	}
}

// Stop stops lsbeat.
func (bt *Lsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func listDir(dirName string, bt *Lsbeat, b *beat.Beat) {
	files, _ := ioutil.ReadDir(dirName)
	fileinfos := []common.MapStr{}

	dirSize := int64(0)
	filecount := 0

	for _, f := range files {
		if f.IsDir() {
			listDir(dirName+"/"+f.Name(), bt, b)
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

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"files":      fileinfos,
			"filesCount": filecount,
			"dirName":    dirName,
			"dirSize":    dirSize,
		},
	}
	bt.client.Publish(event)
}
