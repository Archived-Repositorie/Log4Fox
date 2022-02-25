package log4fox

import "os"

func Stderr(s string) {
	os.Stderr.WriteString(s)
}

func Stdout(s string) {
	os.Stdout.WriteString(s)
}

func Stdin(s string) {
	os.Stdin.WriteString(s)
}