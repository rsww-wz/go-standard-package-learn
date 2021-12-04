package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
POST请求的数据有一下几种
	POST form表单
	POST JSON
	POST 文件

	func Post(url, contentType string, body io.Reader) (resp *Response, err error)
		第一个参数：URL
		第二个参数：提交的数据类型
		第三个参数：从什么地方读取提交的数据


 */

func postForm() {
	/*
		表单数据和查询参数一样，都是需要经过URL编码的
		提交form表单数据的数据类型是：application/x-www-form-urlencoded
		设置好表单之后，对数据进行编码
		第三个参数是io.Reader接口类型，可以选择strings.NewReader
	 */
	formData :=  make(url.Values)
	formData.Add("name","rs")
	formData.Add("age","21")
	form := formData.Encode()
	url := "http://httpbin.org/post"
	r,_ := http.Post(url,"application/x-www-form-urlencoded",strings.NewReader(form))
	defer func() {_ = r.Body.Close()}()
	content,_ :=  ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}


func postJson() {
	/*
		提交JSON数据类型与form表单的主要不同是
			需要结构体转换成JSON数据类型
			1.struct --- JSON（bytes）
			2.bytes.NewReader（JSON）

		通过bytes.NewReader把读入bytes类型的JSON
		JSON数据的类型：application/json

		POST请求的本质
			它是通过request body提交的
			form表单 request body是urlencoding格式
			JSON request body是JSON格式
			文件 request body也是通过组织body数据实现
	 */

	// json数据需要通过和结构体转换
	type Data struct {
		Name string
		age int
	}

	data := Data{
		Name:"rs",
		age:18,
	}

	// 把结构体转换为JSON数据类型
	// 转换后JSON的格式bytes类型
	jsonData,_ := json.Marshal(data)
	url := "http://httpbin.org/post"
	r,_ := http.Post(url,"application/json",bytes.NewReader(jsonData))
	defer func() {_ = r.Body.Close()}()
	content,_ :=  ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}

func main() {
	//postForm()
	postJson()
}
