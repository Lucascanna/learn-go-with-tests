package arrays

import "testing"

func TestArraySum(t *testing.T) {
	t.Run("sum array of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		actual := Sum(numbers)
		expected := 10

		if actual != expected {
			t.Errorf("Expected %q but found %q", expected, actual)
		}
	})
}
