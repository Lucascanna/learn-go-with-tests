package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Test square perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		actual := Perimeter(rectangle)
		expected := 40.0

		if actual != expected {
			t.Errorf("Expected %.2f but gound %.2f", expected, actual)
		}
	})
}
func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("Test rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{7.0, 10.0}
		expected := 70.0

		checkArea(t, rectangle, expected)
	})
	t.Run("Test circle Area", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})
}
