package log4fox

import "testing"

func TestClr(t *testing.T) {
	if FONT_BLACK != "\033[0;30m" {
		t.Error("ColorReset: Invalid constant value")
	} else {
		c := Color(FONT_BLACK, "Hello world!")

		Stderr(c)
	}
}
