package main
import "fmt"
import "sort"

func main() {

	strs := []string { "c", "a", "b" }
	fmt.Println("before sort:", strs)
	sort.Strings(strs)
	fmt.Println("after sort", strs)

	ints := []int { 4, 77, 22, 3423, 999 }

	fmt.Println("before sort:", ints)
	sort.Ints(ints)
	fmt.Println("after sort:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
