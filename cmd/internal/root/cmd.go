package root

import (
	"fmt"
	"github.com/iseki0/cos-uploader/env"
	"github.com/iseki0/cos-uploader/exitcode"
	"github.com/iseki0/cos-uploader/uploader"
	"github.com/spf13/cobra"
	"os"
)

func Cmd() *cobra.Command {
	var c cobra.Command
	c.Use = "cos-upload <FILE...> TARGET"
	c.Args = cobra.MinimumNArgs(2)
	c.Run = run
	c.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
		if e := env.InitEnv(); e != nil {
			fmt.Fprintln(os.Stderr, e.Error())
			exitcode.Set(1)
			exitcode.Exit()
		}
	}
	c.AddCommand(RelayCmd())
	return &c
}

func run(cmd *cobra.Command, args []string) {
	uploader.InitClient()
	var target = args[len(args)-1]
	var files = args[:len(args)-1]
	for _, file := range files {
		e := uploader.UploadFile(file, target)
		if e != nil {
			fmt.Fprintln(os.Stderr, e.Error())
			exitcode.Set(1)
		}
	}
}

func RelayCmd() *cobra.Command {
	var c cobra.Command
	c.Use = "relay"
	c.Run = func(cmd *cobra.Command, args []string) {
		uploader.InitClient()
		e := uploader.Relay()
		if e != nil {
			fmt.Fprintln(os.Stderr, e.Error())
			exitcode.Set(1)
		}
	}
	return &c
}
