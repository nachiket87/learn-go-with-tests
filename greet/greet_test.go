package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nachi")

	got := buffer.String()
	want := "Hello, Nachi"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
