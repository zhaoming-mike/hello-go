package main
import "fmt"

func sum(nums ...int) {
	fmt.Print(nums," ")

	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main(){

	sum(1, 10)
	sum(1,2,3,4,5,7,1000,6)

	nums := []int {1,2,3,4,5,6}
	sum(nums...)
	//cannot use nums (type []int) as type int in argument to sum
	//sum(nums)

}
