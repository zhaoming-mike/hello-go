package main
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file := "/tmp/data-test-for-go-reading-file-del-me"
	file2 := "/tmp/data-test-for-go-reading-file-del-me-2"
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(file, d1, 0644)
	check(err)

	f, err := os.Create(file2)
	check(err)

	defer f.Close()

	d2 := []byte { 115, 111, 109, 65, 10 }
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
