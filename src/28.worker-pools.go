package main
import "fmt"
import "time"

// 请注意channel的“xxx <- chan int” 和 “xxx chan <- int” 写法的区别
func worker(id int, jobs <- chan int, results chan <- int) {
	for j:= range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 9; a++ {
		<- results
	}
}
