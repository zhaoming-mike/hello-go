package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()

	if err != nil {
		panic(err)
	}

	p := fmt.Println
	p("> date")
	p(string(dateOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()

	grepIn.Write([]byte("hello grep\ngoodbye grep"))

	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)

	grepCmd.Wait()
	p("> grep hello")
	p(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	p("> ls -a -l -h")
	p(string(lsOut))
}
