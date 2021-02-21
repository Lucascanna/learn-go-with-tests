package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("word found", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})
	t.Run("word found", func(t *testing.T) {
		_, err := dictionary.Search("pippo")

		want := "Word missing"

		if err == nil {
			t.Fatal("shold return error")
		}

		got := err.Error()
		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "pippo")
		}
	})
}
func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("add", func(t *testing.T) {

		dictionary.Add("pippo", "a dog")

		got, _ := dictionary.Search("pippo")
		want := "a dog"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})
}
