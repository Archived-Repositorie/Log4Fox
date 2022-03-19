package essential

import (
	"bufio"
	"os"
)

type scaningObj struct {
	Get   func() string
	Scan  func() bool
	Error func() error
}

func get(scan *bufio.Scanner) func() string {
	return scan.Text
}

func scan(scan *bufio.Scanner) func() bool {
	return scan.Scan
}

func Create(f *os.File) scaningObj {
	s := bufio.NewScanner(f)
	return scaningObj{
		Scan:  scan(s),
		Get:   get(s),
		Error: errors(s),
	}
}

func errors(scan *bufio.Scanner) func() error {
	return scan.Err
}
