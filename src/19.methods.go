package main
import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	//求面积
	return r.width * r.height
}

func (r rect) perim() int {
	//求周长
	return 2 * r.width + 2 * r.height
}

func main() {
	r := rect{width:10, height:5}

	fmt.Println("area:", r.area())
	fmt.Println("perim", r.perim())


	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
