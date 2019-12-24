package main

import "fmt"

type geometry interface {
	area() float64
	perm() float64
}
type rect struct {
	height, width float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perm() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return 2 * 3.14 * c.radius * c.radius
}
func (c circle) perm() float64 {
	return 2 * 3.14 * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perm())
}
func main() {
	r := rect{24, 4}
	c := circle{6}
	measure(r)
	measure(c)

}
