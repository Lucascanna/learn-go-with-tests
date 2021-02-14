package arrays

import (
	"reflect"
	"testing"
)

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
func TestArraySumAll(t *testing.T) {
	t.Run("sum 3 arrays of any size", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4}
		arr2 := []int{1, 2, 3, 4, 5}
		arr3 := []int{1, 2, 3}
		actual := SumAll(arr1, arr2, arr3)
		expected := []int{10, 15, 6}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v but found %v", expected, actual)
		}
	})
}
func TestArraySumTails(t *testing.T) {
	t.Run("sum 3 tails of any size", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4}
		arr2 := []int{1, 2, 3, 4, 5}
		arr3 := []int{1, 2, 3}
		actual := SumTails(arr1, arr2, arr3)
		expected := []int{9, 14, 5}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v but found %v", expected, actual)
		}
	})
	t.Run("sum 3 tails of any size, one is empty", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4}
		arr2 := []int{}
		arr3 := []int{1, 2, 3}
		actual := SumTails(arr1, arr2, arr3)
		expected := []int{9, 0, 5}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v but found %v", expected, actual)
		}
	})
}
