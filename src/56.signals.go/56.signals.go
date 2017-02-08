package main
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	p := fmt.Println
	go func() {
		sig := <- sigs
		p()
		p("sig:", sig)
		done <- true
	}()

	p("awaiting signal")
	<- done
	p("exiting...")
}
