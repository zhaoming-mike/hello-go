package main
import "fmt"

import "os"

func main() {
	panic("one problem!")

	fmt.Println("os.Create...")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
