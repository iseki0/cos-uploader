package exitcode

import "os"

var code int

func Set(n int) {
	code = n
}

func Exit() {
	os.Exit(code)
}
