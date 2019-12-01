// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"time"
)

type Config struct {
	Period    time.Duration `config:"period"`
	Paths     []string      `config:"paths"`
	SmbDrives []string      `config:"smb_drives"`
}

var (
	DefaultConfig = Config{
		Period: 60 * time.Second,
		Paths:  []string{"."},
	}
)
