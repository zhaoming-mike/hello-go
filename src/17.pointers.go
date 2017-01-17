package main
import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	//值传递
	zeroval(i) 
	fmt.Println("zeroval:", i)

	//引用传递
	zeroptr(&i)	
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)



}
