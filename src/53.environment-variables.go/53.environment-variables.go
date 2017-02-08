package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {

	key := "Foo"
	noValKey := "noValKey"
	os.Setenv(key, "testValue")

	p := fmt.Println

	p("-------------")
	fmt.Printf("%s:%s\n", key, os.Getenv(key))
	fmt.Printf("%s:%s\n", noValKey, os.Getenv(noValKey))

	p()

	for index, e := range os.Environ() {
		pair := strings.Split(e, "=")
		p(index, "key:", pair[0], "=", pair[1])
	}
}
