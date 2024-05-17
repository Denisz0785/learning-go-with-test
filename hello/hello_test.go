package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello", func(t *testing.T) {
		got := Hello("Denis", "")
		want := "Hello, Denis"

		assertMessage(t, got, want)
	})
	t.Run("say Hello World when empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertMessage(t, got, want)
	})
}


func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
