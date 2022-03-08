package color

import (
	"testing"

	"github.com/PurotoApp/Log4Fox/essential"
)

func TestClr(t *testing.T) {
	if FONT_BLACK != "\033[0;30m" {
		t.Error("ColorReset: Invalid constant value")
	} else {
		c := Color(FONT_BLACK, "Hello world!")

		essential.Stderr(c)
	}
}
