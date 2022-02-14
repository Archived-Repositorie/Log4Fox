package main

import "testing"

func TestClr(t *testing.T) {
	if 0 != 1 {
		t.Errorf(string(ColorBFnt[3] + "test"))
	} else {
		t.Logf(string(ColorBFnt[6] + "test"))
	}
}
