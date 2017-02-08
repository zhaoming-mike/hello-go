package main
import "fmt"

func main() {

	nums := []int {1,2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 4 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a":"apple", "b":"jackson", "c":"michael"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i,c:= range "ABCDE" {
		fmt.Println(i,c)
	}

}
