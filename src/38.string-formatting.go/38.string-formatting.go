package main

import "fmt"
import "os"

var pf = fmt.Printf

type point struct {
	x, y int
}

func main() {
	p := point { 1, 2 }
	fmt.Println("-----------------------")
	fmt.Printf("%v\n", p)					//只打印字段值
	fmt.Printf("%+v\n", p)					//字段名称、字段（类似 json 字符串）
	pf("%#v\n", p)							//类型、字段名称、字段值
	pf("%T\n", p)							//只打印类型
	pf("%t\n", true)
	pf("%d\n", 1234)
	pf("%b\n", 8)							//二进制
	pf("%c\n", 65)							//printf by code (ASCII)
	pf("%x\n", 16)							//十六进制编码
	pf("%f\n", 78.9)
	pf("%e\n", 123400000.0)					//科学计数法小写e形式
	pf("%E\n", 123400000.0)					//科学计数法大写E形式

	pf("%s\n", "\"michael\"")				//转义输出
	pf("%q\n", "\"michael\"")				//原样输出

	pf("%x\n", "hex this")					//base-16
	pf("%p\n", &p)

	pf("|%10d|%10d|\n", 12, 1234567890)		//制表打印
	pf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	pf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)		//左对齐

	pf("|%6s|%6s|\n", "hello", "sdfsdfs")

	s := fmt.Sprintf("a %s", "String")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
	fmt.Fprintf(os.Stdout, "an %s\n", "out")



}
