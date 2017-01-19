package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var p = fmt.Println
func main() {
	p("---------------")
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	p(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	p(r.MatchString("pppppch"))

	p(r.FindString("punch peach pppch"))
	p(r.FindStringIndex("punch peach pppch"))		//第一个找到的字符串的开始和结尾的字符索引
	p(r.FindStringSubmatch("punch peach pppch"))	//返回包括完整表达式匹配内容以及小括号中的匹配内容
	p(r.FindStringSubmatchIndex("punch peach pppch"))

	p(r.FindAllString("peach xxx punch pinch", -1))					//匹配所有内容
	p(r.FindAllString("peach xxx punch pinch", 2))					//匹配n个内容
	p(r.FindAllStringSubmatchIndex("peach xxx punch pinch", 2))		//匹配n个内容

	p(r.Match([]byte("peach")))										//二进制也支持

	r = regexp.MustCompile("p([a-z]+)ch")
	p(r)
	p(r.ReplaceAllString("a peach is not very good peach", "<fruit>"))

	in := []byte("a peach is very good peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)						//这里的Func指的是什么？
	p(string(out))
}
