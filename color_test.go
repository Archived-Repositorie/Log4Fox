package github.com/PurotoApp/Log4Fox

import "testing"

func TestClr(t *testing.T) {
	if ColorReset[0] != "\033[m" {
		t.Error("ColorReset: Invalid constant value")
	}
}