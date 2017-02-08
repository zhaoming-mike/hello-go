package main
import "fmt"
import "math/rand"

func main() {
	p := fmt.Println
	p("--------------------")
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	p()

	p(rand.Float64())

	fmt.Print((rand.Float64() * 5) + 5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	p()

	s1 := rand.NewSource(42)
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	p()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	p()
}
