package main
import "os"
import "fmt"

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	p := fmt.Println

	p(argsWithProg)
	p(argsWithoutProg)
	p(arg)

	//go run xxx a b c d
}
