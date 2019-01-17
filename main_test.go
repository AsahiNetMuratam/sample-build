package main

import (
	"testing"
)

func TestMessage(t *testing.T) {
	if Message() != "sample" {
		t.Fatal()
	}
}
