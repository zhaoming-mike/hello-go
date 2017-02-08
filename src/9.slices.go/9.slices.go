package main
import "fmt"

func main() {

	s := make([]string , 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get1", s[1])
	fmt.Println("get2", s[2])
	fmt.Println("length:", len(s))
	
	s = append(s, " mike")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	c := make([] string, len(s) )
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:6]
	fmt.Println("slice 1:", l)

	l = s[:5]
	fmt.Println("slice 2:", l)

	l = s[2:]
	fmt.Println("slice 3:", l)

	t := [] string {"michael", "jackson", "mike"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i:=0;i<3;i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j:=0;j<innerLen;j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)








}
