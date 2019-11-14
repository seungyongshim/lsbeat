package main

import (
	"os"

	"github.com/seungyongshim/lsbeat/cmd"

	_ "github.com/seungyongshim/lsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
