package structures

import (
	"math"
)

type Rectangle struct {
	width float64
	hight float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base  float64
	hight float64
}

type Shape interface {
	Area() float64
}

func Periemetr(r Rectangle) float64 {
	return (r.width + r.hight) * 2
}

func (r *Rectangle) Area() float64 {
	return r.width * r.hight
}

func (c *Circle) Area() float64 {
	return (c.radius * c.radius) * math.Pi
}

func (t *Triangle) Area() float64 {
	return (t.base * t.hight) * 0.5
}
