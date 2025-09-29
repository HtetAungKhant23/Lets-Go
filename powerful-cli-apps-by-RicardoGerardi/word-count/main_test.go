package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("testing one million two million three million")
	expCount := 7
	res := count(b, false)
	if res != expCount {
		t.Errorf("Expected %d, got %d instead.\n", expCount, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("testing\n one million\n two million\n three million")
	expCount := 4
	res := count(b, true)
	if res != expCount {
		t.Errorf("Expected %d, got %d instead.\n", expCount, res)
	}
}