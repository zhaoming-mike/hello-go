package main
import "strconv"
import "fmt"

func main() {
	p := fmt.Println
	f, _ := strconv.ParseFloat("1.234", 64)
	p("f:", f)

	i, _ := strconv.ParseInt("123", 0, 64)
	p("i:", i)

	d, _ := strconv.ParseUint("789", 0, 64)
	p("d:", d)

	k, _ := strconv.Atoi("135")
	p("k", k)

	_, e := strconv.Atoi("sdfsf")
	p("err:", e)
}
