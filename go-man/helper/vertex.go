package helper

import "math"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Dist(p Vertex) float64 {
	x := p.X - v.X
	y := p.Y - v.Y
	return math.Sqrt(x * x + y * y)
}
