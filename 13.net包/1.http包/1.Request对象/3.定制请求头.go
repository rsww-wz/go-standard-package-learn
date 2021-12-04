package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 创建请求
	request,err := http.NewRequest(http.MethodGet,"http://www.baidu.com",nil)
	if err != nil {
		fmt.Println("Get error", err)
		return
	}

	/*
	设置请求头
		请求头也是一个字典
		添加：func (h Header) Add(key, value string)
		设置：func (h Header) Set(key, value string)
		删除：func (h Header) Del(key string)
		是否有：func (h Header) has(key string) bool
		获取value（key忽略大小写）
			func (h Header) Get(key string) string
		获取value（key不忽略大小写）
			func (h Header) get(key string) string
	 */
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.134 Safari/537.36")
	request.Header.Add("origin", "https://www.zhihu.com")
	request.Header.Add("referer","https://www.zhihu.com/")

	request.Header.Set("origin","http://www.baidu.com")
	request.Header.Del("referer")

	s1 := request.Header.Get("user-agent")
	s2 := request.Header.Get("User-Agent")
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := request.Header.Get("origin")
	s4 := request.Header.Get("referer")
	fmt.Println("origin的value是：",s3)
	fmt.Println("referer的value还存在吗：",s4)
}
