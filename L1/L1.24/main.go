package main

import "math"

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(q *Point) float64 {
	if p == nil || q == nil {
		return 0
	} else {
		return float64(math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y))))
	}
}

func main() {
	a := NewPoint(1, 2)
	b := NewPoint(3, 4)

	println(a.Distance(b))
}
