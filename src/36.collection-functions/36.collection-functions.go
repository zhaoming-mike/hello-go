package main
import "strings"
import "fmt"

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string { "peach", "apple", "pear", "plum" }
	fmt.Println("strs:", strs)
	fmt.Println("find `pear` index of strs is:", Index(strs, "pear"))

	fmt.Println("strs is find `grape`?:", Include(strs, "grape"))

	fmt.Println("is has prefix `p` in strs:", Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println("is not has prefix `p` in strs:", All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println("contains `e` append together:", Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println("strs in upper case:", Map(strs, strings.ToUpper))
}
