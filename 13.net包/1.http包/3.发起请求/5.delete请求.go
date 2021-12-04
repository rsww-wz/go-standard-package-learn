package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//和put方法一样，系统也直接提供delete方法，都是需要通过client手动发起请求

func main() {
	delete()
}

func delete() {
	request,err := http.NewRequest(http.MethodDelete,"http://httpbin.org/delete",nil)
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
