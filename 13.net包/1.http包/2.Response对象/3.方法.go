package main

func main() {

}

/*
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
	ReadResponse从r读取并返回一个HTTP 回复
	req参数是可选的，指定该回复对应的请求（即是对该请求的回复）
	如果是nil，将假设请求是GET请求。客户端必须在结束resp.Body的读取后关闭它
	读取完毕并关闭后，客户端可以检查resp.Trailer字段获取回复的trailer的键值对

func (r *Response) Cookies() []*Cookie
	Cookies解析并返回该回复中的Set-Cookie头设置的cookie

func (r *Response) Location() (*url.URL, error)
	Location返回该回复的Location头设置的URL
	相对地址的重定向会相对于该回复对应的请求来确定绝对地址
	如果回复中没有Location头，会返回nil, ErrNoLocation
 */
