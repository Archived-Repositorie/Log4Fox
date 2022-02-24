package log4fox

import "testing"

func TestClr(t *testing.T) {
	if Color().Font.BLACK != "\033[0;30m" {
		t.Error("ColorReset: Invalid constant value")
	}
}
