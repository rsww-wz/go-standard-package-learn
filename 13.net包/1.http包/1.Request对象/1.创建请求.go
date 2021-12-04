package main

import (
	"fmt"
	"net/http"
)

/*
创建请求：func NewRequest(method, url string, body io.Reader) (*Request, error)
	第一个参数：请求方法
	第二个参数：请求地址
	第三个参数：

	返回值：*Request, error

	如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body
	并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭
 */

func main() {
	//func NewRequest(method, url string, body io.Reader) (*Request, error)
	request,err := http.NewRequest(http.MethodGet,"http://www.baidu.com",nil)
	if err != nil {
		fmt.Println("Get error", err)
		return
	}

	fmt.Printf("%s：%T\n","request的类型是",request)
	fmt.Printf("%s：%v\n","request的值是",request)
}
