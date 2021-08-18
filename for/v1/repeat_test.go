package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("repeated %q expected %q", repeated, expected)
		}
	}
	t.Run("Repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("Repeat 3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
