package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	rdr := bufio.NewReader(os.Stdin)
	out := os.Stdout

	for {
		switch line, err := rdr.ReadString('\n'); err {
		case nil:
			ucl := strings.ToUpper(line)
			if _, err = out.WriteString(ucl); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
}
