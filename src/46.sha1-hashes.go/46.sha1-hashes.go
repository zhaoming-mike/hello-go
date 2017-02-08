package main
import "crypto/sha1"
import "fmt"

var p = fmt.Println

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	p("s:", s)
	fmt.Printf("%x\n", bs)
}
