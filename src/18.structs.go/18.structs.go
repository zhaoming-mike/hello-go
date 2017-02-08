package main
import "fmt"

type person struct {
	id int
	name string
	age	int
}


func main() {

	fmt.Println(person{1, "Mike", 34})
	fmt.Println(person{name:"michael", id:1, age:34})

	//未初始化的值就是zero-valued.
	fmt.Println(person{age:50})

	fmt.Println(&person{id:100, name:"jackson", age:52})

	s := person{name:"Sean", age:99}
	fmt.Println("s.name:", s.name)
	fmt.Println("s.id:", s.id)
	fmt.Println("s.age:", s.age)
}
