package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
post函数提交表单需要对表单进行URL编码
而PostForm函数可以把表单写入参数，由函数对表单进行编码，简化工作
由于提交的是表单，也就是不需要指定提交的数据类型了

func PostForm(url string, data url.Values) (resp *Response, err error)

 */

func main() {
	// 初始化字典
	data := make(url.Values)

	// 设置键值对
	// map[string][]string
	// value是字符串切片，所以需要用花括号
	data = url.Values{
		"name":{"rs"},
		"age":{"21"},
	}

	// 添加键值对
	data.Add("id"," CO006635")

	r,_ := http.PostForm("http://httpbin.org/post",data)
	defer func() {_ = r.Body.Close()}()
	content,_ :=  ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}
