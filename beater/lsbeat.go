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
	path          string    // root 디렉토리
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

	bt.path = c.Path

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
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		listDir(bt.path, bt, b, counter)

		bt.lastIndexTime = time.Now()
		logp.Info("Event sent")
		counter++
	}
}

// Stop stops lsbeat.
func (bt *Lsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func listDir(dirFile string, bt *Lsbeat, b *beat.Beat, counter int) {
	files, _ := ioutil.ReadDir(dirFile)
	for _, f := range files {
		t := f.ModTime()

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":        b.Info.Name,
				"counter":     counter,
				"modTime":     t,
				"filename":    f.Name(),
				"fullname":    dirFile + "/" + f.Name(),
				"isDirectory": f.IsDir(),
				"fileSize":    f.Size(),
			},
		}

		// 첫번째 실행인 경우 전체 파일 목록 색인
		if counter == 1 {
			bt.client.Publish(event)
		} else {
			// 2번째 이후 인 경우 실행 이후 추가된 파일만 색인
			if t.After(bt.lastIndexTime) {
				bt.client.Publish(event)
			}
		}

		// 디렉토리인 경우 하위 파일들 재귀 호출.
		if f.IsDir() {
			listDir(dirFile+"/"+f.Name(), bt, b, counter)
		}
	}
}
