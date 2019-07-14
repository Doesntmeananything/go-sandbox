package structs

import "math"

// Shape describes a type with an Area() method.
type Shape interface {
	Area() float64
}

// Rectangle describes a rectangular shape.
type Rectangle struct {
	Width  float64
	Height float64
}

// Triangle describes a triangular shape.
type Triangle struct {
	Base   float64
	Height float64
}

// Circle describes a circular shape.
type Circle struct {
	Radius float64
}

// Perimeter takes in width and height of a rectangular shape and returns its perimeter.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area for a rectangular shape.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area for a triangular shape.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Area for a circular shape.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
