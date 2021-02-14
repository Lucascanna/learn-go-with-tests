package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat single character", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"

		if actual != expected {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	})
	t.Run("Repeat double character", func(t *testing.T) {
		actual := Repeat("bb", 4)
		expected := "bbbbbbbb"

		if actual != expected {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	})
}
