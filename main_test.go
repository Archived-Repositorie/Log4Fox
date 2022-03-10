package log4fox

import (
	"testing"
)

var testLogger = loggerObject{Name: "test", Color: true}

func TestLog(t *testing.T) {
	if err := testLogger.Logln("Test"); err != nil {
		t.Error(err)
	}
}
