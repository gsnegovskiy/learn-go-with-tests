package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d, %v", got, want, numbers)
		}
	}
	t.Run("Repeat any times", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertCorrectMessage(t, got, want, numbers)
	})
}
func TestSumAll(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want []int, allArgs ...[]int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d, input %v", got, want, allArgs)
		}
	}
	t.Run("Repeat any times", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertCorrectMessage(t, got, want, []int{1, 2}, []int{0, 9})
	})
}
