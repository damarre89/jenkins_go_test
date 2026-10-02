package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("trigger test", func(t *testing.T) {
		got := Greet()
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
