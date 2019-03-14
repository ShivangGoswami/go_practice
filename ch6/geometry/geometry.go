package main

import "math"

type Point struct {
	X, Y float64
}

//Distance :traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Distance : thing but as method of the point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// var a = func(p, q Point) float64 {
// 	return math.Hypot(q.X-p.X, q.Y-p.Y)
// }

// var b=func (p Point) (q Point) float64 {
// 	return math.Hypot(q.X-p.X, q.Y-p.Y)
// }
