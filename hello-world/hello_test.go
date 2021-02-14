package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Luca", "")
		want := "Hello, Luca"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello to world if empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello to people in Spanish", func(t *testing.T) {
		got := Hello("Luca", "spanish")
		want := "Hola, Luca"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello to people in French", func(t *testing.T) {
		got := Hello("Luca", "french")
		want := "Bonjour, Luca"
		assertCorrectMessage(t, got, want)
	})
}
