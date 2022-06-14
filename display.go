package main

import (
	"fmt"
	"github.com/muesli/termenv"
)

var _info = termenv.String("INFO").Foreground(termenv.ANSIBrightWhite)
var _warn = termenv.String("WARN").Foreground(termenv.ANSIBrightYellow)
var _error = termenv.String("ERROR").Foreground(termenv.ANSIBrightRed)

func PrintInfo(a ...any) {
	fmt.Printf("[%s] %s\n", _info, fmt.Sprint(a...))
}

func PrintWarn(a ...any) {
	fmt.Printf("[%s] %s\n", _warn, fmt.Sprint(a...))
}

func PrintError(a ...any) {
	fmt.Printf("[%s] %s\n", _error, fmt.Sprint(a...))
}
