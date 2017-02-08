package main
import (
	"fmt"
	"time"
	"math"
	"strconv"
)


func main() {

	fmt.Println("time.Second:", time.Second)
	fmt.Println("time.Second(times 60):", time.Second * 60)
	fmt.Println("math.Pi:", math.Pi)
	queue := make(chan string, 17)
	queue <- "one"
	queue <- "two"
	
	for i := 0; i < 15; i++ {
		queue <- "hello, " + strconv.Itoa(i)
	}

	close(queue)

	for elem := range queue {
		fmt.Println("elem", elem)
	}
}

