package main

import (
	"encoding/json"
	"fmt"
)

// 序列化: 结构，字典转换成json字符串
// 注意:如果要转换成json字符串，结构体以及属性对外必须是可见的
// 因为序列化需要用json包下的api，这些结构体需要对外可见才能被json包下的api读取到
type Person struct {
	Name  string
	Age   int
	Rmb   float64
	Sex   bool
	Hobby []string
}

func main() {
	person := Person{
		Name:  "于谦",
		Age:   50,
		Rmb:   123.45,
		Sex:   true,
		Hobby: []string{"抽烟", "喝酒", "烫头"},
	}

	//func Marshal(v interface{}) ([]byte, error)
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
