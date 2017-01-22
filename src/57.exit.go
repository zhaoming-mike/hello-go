package main

import "fmt"
import "os"

func main() {
	fmt.Println("-----------")
	defer fmt.Println("!")
	os.Exit(3)
}
