package main
import "fmt"
import "strconv"
func main() {

	s := make([]string , 3)
	fmt.Println("emp:", s)
	fmt.Println("len:", len(s))


	for i:=0; i<3; i++ {
		s[i] = "name-" + strconv.Itoa(i)
	}

	fmt.Println("slice context:", s)

	vvv := "123"
	fmt.Println("vvv:", 1 + rune(vvv))

}
