package main

import (
	"os"

	"github.com/elastic/beats/libbeat/version"

	"github.com/seungyongshim/lsbeat/cmd"

	_ "github.com/seungyongshim/lsbeat/include"
)

var qualifier string

func main() {
	version.SetQualifier(qualifier)

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
