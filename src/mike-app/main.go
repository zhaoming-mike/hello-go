package main

import "hello-go/src/mike-app/lib1" //这里引号内的是文件路径
import "fmt"

func main() {
	p := fmt.Println
	lib.User()	//这里lib是真正的导入的路径下的package名称
	p("end")


	/*
	import "xx1/xx2/xx3"

	func xxx() {
		xx3.xxx()	
		//这里的xx3在import关键词后面的是文件路径
		//在方法xxx()前的是包名，原则上路径跟包名保持一致，但可以不同。
	}	
	
	*/
}
