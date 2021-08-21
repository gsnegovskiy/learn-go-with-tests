package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Triangle) Area() float64 {
	return r.Height * r.Width / 2
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Circle) Area() float64 {
	return math.Pow(r.Radius, 2) * math.Pi
}
