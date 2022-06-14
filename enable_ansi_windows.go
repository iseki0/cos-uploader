//go:build windows

package main

import "github.com/muesli/termenv"

func init() {
	mustA(termenv.EnableWindowsANSIConsole())
}
