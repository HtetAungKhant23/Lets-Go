package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("testing one million two million three million\n")
	expCount := 7
	res := count(b)
	if res != expCount {
		t.Errorf("Expected %d, got %d instead.\n", expCount, res)
	}
}