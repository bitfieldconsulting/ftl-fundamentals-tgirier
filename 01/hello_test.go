package hello

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	want := "Hello Gophers!"
	got := Greeting()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
