package log4fox

import (
	"bufio"
	"os"
)

func Stderr(s string) error {
	_, err := os.Stderr.WriteString(s)
	return err
}

func Stdout(s string) error {
	_, err := os.Stdout.WriteString(s)
	return err
}

func Stdin(s *string) error {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*s = scanner.Text()
	}
	return scanner.Err()
}