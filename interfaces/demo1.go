package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}
type rectangles struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangles) area() float64 {
	return r.width * r.height
}

func school(s shape) {
	fmt.Println("Şeklin alanı: ", s.area())

}
func Demo1() {

	r := rectangles{width: 5, height: 10}
	school(r)
	c := circle{3}
	school(c)

}
