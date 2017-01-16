package main
import "fmt"

func main() {

	m := make(map[string]int)
	fmt.Println("maps:", m)

	m["key1"] = 100
	m["key2"] = 20000


	fmt.Println("maps:", m)
}
