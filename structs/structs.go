package structs

import "math"

// Perimeter function
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Heigth)
}

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle type
type Rectangle struct {
	Width  float64
	Heigth float64
}

// Area function
func (r Rectangle) Area() float64 {
	return r.Heigth * r.Width
}

// Circle type
type Circle struct {
	Radius float64
}

// Area function
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
