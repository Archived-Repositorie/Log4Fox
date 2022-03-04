package log4fox

import "testing"

func TestClr(t *testing.T) {
	if C.Font.BLACK != "\033[0;30m" {
		t.Error("ColorReset: Invalid constant value")
	} else {
		c := Color(C.Font.BLACK, "Hello world!")
		
		Stderr(c)
	}
}
