// Sfangfa
package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello World!")
	var r Rect
	r.max = Point{2, 2}
	fmt.Println(r.Area())
}

type Point struct {
	x, y float64
}

type Rect struct {
	min, max Point
}

func (r *Rect) Area() float64 {
	w := r.max.x - r.min.x
	h := r.max.y - r.min.y
	return w * h
}
