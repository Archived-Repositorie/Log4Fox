package log4fox

import (
	"testing"

	c "github.com/PurotoApp/Log4Fox/color"
	"github.com/PurotoApp/Log4Fox/essential"
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

func TestLoc(t *testing.T) {
	if file, _, line, ok := essential.Locate(1); !(file == "main_test.go" && line == 29 && ok == true) {
		t.Errorf("Locate not correct, returned values: file:%v line:%v ok:%v", file, line, ok)
	}
}