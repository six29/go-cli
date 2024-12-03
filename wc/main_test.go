package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected := 4
	got := count(b, false)

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2\nword3 word4\nword5\n")
	expected := 3
	got := count(b, true)

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}
}
