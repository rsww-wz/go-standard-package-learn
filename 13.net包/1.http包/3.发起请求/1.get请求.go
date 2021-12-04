package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	其实http包下有可以直接发起get，POST，delete请求的函数，但是不能管理请求头等参数
	通过新建请求，设置请求参数，然后通过Client发起请求的方法，可以自定义需要发起请求的内容

	func Get(url string) (resp *Response, err error)
 */

func get() {
	r,err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer func() {_= r.Body.Close()}()

	content,err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func main () {
	get()
}
