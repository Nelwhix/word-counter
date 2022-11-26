package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("I am Nelson, testing my word counter\n")
	expected := 7
	actual := count(b, false, false)

	if expected != actual {
		t.Errorf("Expected %d, got %d instead.\n", expected, actual)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("I am Nelson\n testing my word counter\n")
	expected := 2
	actual := count(b, true, false)

	if expected != actual {
		t.Errorf("Expected %d, got %d instead.\n", expected, actual)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("I am Nelson\n testing my word counter\n")
	expected := 37
	actual := count(b, false, true)

	if expected != actual {
		t.Errorf("Expected %d, got %d instead.\n", expected, actual)
	}
}