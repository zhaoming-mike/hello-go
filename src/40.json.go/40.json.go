package main
import "encoding/json"
import "fmt"
import "os"

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

var p = fmt.Println

func main() {
	p("-------------------")
	bolB, _ := json.Marshal(true)
	p("bool:", string(bolB))

	intB, _ := json.Marshal(1)
	p("int:", string(intB))

	fltB, _ := json.Marshal(2.34)
	p("float:", string(fltB))

	strB, _ := json.Marshal("hello michael")
	p("string:", string(strB))

	slcD := []string { "apple", "peach", "pear" }
	slcB, _ := json.Marshal(slcD)

	p("string array:", string(slcB))

	mapD := map[string]int { "apple":5, "lettuce":7 }
	mapB, _ := json.Marshal(mapD)
	p("map:", string(mapB))

	res1D := &Response1 { Page:1, Fruits:[]string { "a", "b", "v" } }
	res1B, _ := json.Marshal(res1D)
	p("struct:", string(res1B))

	res2D := &Response2 { Page:1, Fruits:[]string { "a1", "b1", "c1" } }
	res2B, _ := json.Marshal(res2D)
	p("struct2:", string(res2B))

	byt := []byte(`{"num":6.0, "strs":["zzz", "bbb"]}`)
	var dat map[string]interface {}

	//反序列化成制定的对象，如果错误就打印信息
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	p("dat:", dat)

	num := dat["num"].(float64)
	p("num:", num)

	strs := dat["strs"].([]interface{})		//这里的kv访问方式要特别注意
	str1 := strs[0].(string)
	p("str1:", str1)

	str := `{"page":1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	//反序列化成结构体
	json.Unmarshal([]byte(str), &res)
	p("res:", res)
	p("res.Fruits[0]:", res.Fruits[0])

	//编码并打印到标准输出流
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int { "apple":100, "peach":90 }
	enc.Encode(d)
}
