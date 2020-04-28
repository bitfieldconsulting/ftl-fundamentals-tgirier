package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Howdy, Gopherinos!"
	got := hello.Greeting()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
