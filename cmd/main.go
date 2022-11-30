package main

import (
	"fmt"
	"github.com/iseki0/cos-uploader/cmd/internal/root"
	"github.com/iseki0/cos-uploader/exitcode"
	"os"
)

func main() {
	if e := root.Cmd().Execute(); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(1)
	}
	exitcode.Exit()
}
