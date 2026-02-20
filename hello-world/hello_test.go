package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Amogh")
	want := "Hello, Amogh"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
