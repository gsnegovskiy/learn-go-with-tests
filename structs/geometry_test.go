package structs

import "testing"

func TestPerimeter(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("Calculate rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Rectangle.Perimeter(rectangle)
		want := 40.0
		assertCorrectMessage(t, got, want)
	})
}

func TestArea(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertCorrectMessage(t, got, want)
	}

	t.Run("Calculate rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})
	t.Run("Calculate circle area", func(t *testing.T) {
		circle := Circle{100.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
