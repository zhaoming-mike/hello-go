package main
import "fmt"

func main() {

	m := make(map[string]int)
	fmt.Println("maps:", m)

	m["key1"] = 100
	m["key2"] = 20000
	m["mike"] = 999

	fmt.Println("maps:", m)

	v1 := m["mike"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "mike")
	fmt.Println("len:", len(m))
	fmt.Println("maps:", m)

	prs := m["mike"]

	fmt.Println("prs:", prs)



}
