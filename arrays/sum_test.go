package arrays

import "testing"

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want int, numbers [5]int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d, %v", got, want, numbers)
		}
	}
	t.Run("Repeat 5 times", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want, numbers)
	})
}
