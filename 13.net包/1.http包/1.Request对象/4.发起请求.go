package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// creatRequest 创建请求并发送请求
func creatRequest (url string) *http.Response {
	request,err := http.NewRequest(http.MethodGet,url,nil)
	if err != nil {
		fmt.Println("Get error", err)
	}

	// 设置请求头
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.134 Safari/537.36")

	// 发送请求
	//func (c *Client) Do(req *Request) (*Response, error)
	r,err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	return r
}

// 并打印响应返回值
func printBody(r *http.Response) {
	content,err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	fmt.Println(string(content))
}

func main() {
	url := "http://www.baidu.com"
	content := creatRequest(url)
	printBody(content)
}
