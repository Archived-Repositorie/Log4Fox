package log4fox

import (
	"bufio"
	"os"
)

func Stderr(s string) {
	os.Stderr.WriteString(s)
}

func Stdout(s string) {
	os.Stdout.WriteString(s)
}

func Stdin(s *string) error {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*s = scanner.Text()
	}
	return scanner.Err()
}