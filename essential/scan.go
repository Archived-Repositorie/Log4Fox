package essential

import (
	"bufio"
	"os"
)

func ReaderStr(s *bufio.Scanner,r bool) string {
	if r {
		return s.Text()
	} else {
		return ""
	}
}

func Reader(s *os.File) (bool, *bufio.Scanner) {
	scan := bufio.NewScanner(s)
	return scan.Scan(), scan
}

func ReaderErr(s *bufio.Scanner) error {
	return s.Err()
}
