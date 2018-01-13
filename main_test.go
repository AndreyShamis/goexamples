package main

import "testing"

func TestStart(t *testing.T) {
	start := printStart()
	if start != "Starting application." {
		t.Errorf("Bad start string %s", start)
	}
}
