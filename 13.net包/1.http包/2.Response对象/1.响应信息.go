package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 发送请求
func sendRequest (url string) *http.Response {
	request,err := http.NewRequest(http.MethodGet,url,nil)
	if err != nil {
		fmt.Println("Get error", err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.134 Safari/537.36")
	r,err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	return r
}

// 获取响应主体
func printBody(r *http.Response) {
	content,err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

// 获取响应状态信息
func printInfo(r *http.Response) {
	// 获取状态码
	fmt.Println(r.StatusCode)

	// 获取状态码信息
	fmt.Println(r.Status)

	// 使用的http协议
	fmt.Println(r.Proto)

	// 其他信息
	fmt.Println(r.ProtoMajor)
	fmt.Println(r.ProtoMinor)
}

// 获取响应头
func printHeader(r *http.Response) {
	// 获取用内容类型，可以忽略大小写
	fmt.Println(r.Header.Get("Content-type"))
}
func main() {
	url := "http://baidu.com"
	r := sendRequest(url)
	defer r.Body.Close()
	printHeader(r)
	printInfo(r)
	printBody(r)

}
