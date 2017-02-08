package main
import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perim() float64 {
	return 2 * s.width + 2 * s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("geometry:", g)
	fmt.Println("geometry.area():", g.area())
	fmt.Println("geometry.perim()", g.perim())
}

func main(){
	s := square{width:3, height:4}
	c := circle{radius:5}
	measure(s)
	measure(c)
}
