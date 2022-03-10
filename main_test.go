package log4fox

import (
	"testing"

	"github.com/PurotoApp/Log4Fox/essential"
	c "github.com/PurotoApp/Log4Fox/color"
)

var testLogger = loggerObject{Name: "test", Color: true}

func TestLog(t *testing.T) {
	if err := testLogger.Logln("Test"); err != nil {
		t.Error(err)
	}
}

func TestClr(t *testing.T) {
	if c.FONT_BLACK != "\033[0;30m" {
		t.Error("ColorReset: Invalid constant value")
	} else {
		c := c.Color(c.FONT_RED_UL, "Hello world!")

		essential.Stderr(c)
	}
}