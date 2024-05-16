package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Denis")
	want := "Hello, Denis"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}
