package main

import (
	"testing"
	"log"
)

func TestStart(t *testing.T) {
	start := printStart()
	if start != "Starting application." {
		t.Errorf("Bad start string %s", start)
	}

	out, err := getIfconfig()
	log.Printf("out: %s , err : %s", out, err)
}
