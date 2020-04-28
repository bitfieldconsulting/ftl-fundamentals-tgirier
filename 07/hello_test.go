package hello_test

import (
	"hello"
	"testing"
)

var greetingTests = []string{
	"Hi there",
	"Hello",
}

func TestReturnGreeting(t *testing.T) {
	for _, n := range greetingTests {
		want := n + " yourself!"
		got := hello.ReturnGreeting(n)
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	}
}
