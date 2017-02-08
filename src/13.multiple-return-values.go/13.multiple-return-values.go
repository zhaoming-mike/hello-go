package main
import "fmt"

func vals() (int, int) {
	return 1000,9000
}

func main(){
	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, e := vals()
	fmt.Println("e:>>>", e)
}
