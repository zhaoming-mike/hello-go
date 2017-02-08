package main
import "fmt"
import "sort"

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string { "peach", "banana", "apple", "kiwi", "orange"}

	sort.Sort(ByLength(fruits))

	fmt.Println("sort by word lenth fruits:", fruits)

}
