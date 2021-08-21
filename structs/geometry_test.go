package structs

import "testing"

func TestPerimeter(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	assertCorrectMessage(t, got, want)

}

func TestArea(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	want := 100.0
	assertCorrectMessage(t, got, want)

}
