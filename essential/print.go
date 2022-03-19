package essential

import (
	"golang.org/x/sys/unix"
)

func Stderr(s string) error {
	_, err := unix.Write(unix.Stderr, []byte(s))
	return err
}

func Stdout(s string) error {
	_, err := unix.Write(unix.Stdout, []byte(s))
	return err
}

