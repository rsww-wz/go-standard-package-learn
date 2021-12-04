package main

import (
	"fmt"
	"net/http"
	"net/url"
)


func main() {
	// 创建请求
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}

	/*
	设置URL参数
		type Values map[string][]string
		values类型是一个字典，所以需要用make函数初始化
		字典的方法
			func (v Values) Get(key string) string
			func (v Values) Set(key, value string)
			func (v Values) Add(key, value string)
			func (v Values) Del(key string)
	*/
	params := make(url.Values)
	params.Add("name", "rs")
	params.Add("age", "18")
	params.Add("gender","male")

	params.Set("age","21")
	params.Del("gender")

	value1 := params.Get("age")
	value2 := params.Get("gender")
	fmt.Println("age的value是：",value1)
	fmt.Println("gender的value是：",value2)

	// 参数绑定URL地址
	// URL是一个结构体，RawQuery是结构体的一个属性
	// Encode()是params对象的一个方法：作用就是将参数编码为URL编码格式
	request.URL.RawQuery = params.Encode()

	// 打印编码成URL格式后的参数
	fmt.Println(params.Encode())   //age=18&name=rs
}
