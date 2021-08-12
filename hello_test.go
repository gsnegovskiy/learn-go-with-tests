package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Greg", "")
		want := "Hello, Greg"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Greg", "Spanish")
		want := "Hola, Greg"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Greg", "French")
		want := "Bonjour, Greg"
		assertCorrectMessage(t, got, want)
	})

}
