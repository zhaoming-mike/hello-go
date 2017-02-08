package main
import (
	"syscall"
	"os"
	"os/exec"
)

func main() {
	binary, err := exec.LookPath("ls")

	if err != nil {
		panic(err)
	}

	args := []string{ "-a", "-l", "-h" }

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}
}
