package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
go语言没有直接提供发起put请求的函数
但是可以通过创建put请求，然后通过client发起请求
 */

func main() {
	put()
}

func put() {
	request,err := http.NewRequest(http.MethodPut,"http://httpbin.org/put",nil)
	if err != nil {
		panic(err)
	}
	r,err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()
	content,err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
