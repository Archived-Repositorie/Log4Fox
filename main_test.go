package log4fox

import (
	"testing"
)

var testLogger = loggerObject{name: "test"}

func TestLog(t *testing.T) {
	if err :=testLogger.Logln("Test"); err != nil {
		t.Error(err)
	}
}
