package main
import "fmt"
import "time"

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	<- timer1.C

	fmt.Println("Timer 1 expired")

	//timer2 := time.NewTimer(time.Second * 5)
	timer2 := time.NewTimer(time.Second)
	go func() {
		<- timer2.C
		fmt.Println("Timer 2 expired")
	}()

	//stop2 := false
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
