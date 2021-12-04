package main

func main() {

}

/*
func (c *Client) Do(req *Request) (resp *Response, err error)
	Do方法发送请求，返回HTTP回复,它会遵守客户端c设置的策略（如重定向、cookie、认证）
	如果客户端的策略（如重定向）返回错误或存在HTTP协议错误时，本方法将返回该错误；如果回应的状态码不是2xx，本方法并不会返回错误。
	如果返回值err为nil，resp.Body总是非nil的，调用者应该在读取完resp.Body后关闭它
	如果返回值resp的主体未关闭，c下层的RoundTripper接口（一般为Transport类型）可能无法重用resp主体下层保持的TCP连接去执行之后的请求。
	请求的主体，如果非nil，会在执行后被c.Transport关闭，即使出现错误
	一般应使用Get、Post或PostForm方法代替Do方法

func (c *Client) Head(url string) (resp *Response, err error)
	Head向指定的URL发出一个HEAD请求，如果回应的状态码如下，Head会在调用c.CheckRedirect后执行重定向

func (c *Client) Get(url string) (resp *Response, err error)
	Get向指定的URL发出一个GET请求，如果回应的状态码如下，Get会在调用c.CheckRedirect后执行重定向：

func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
	Post向指定的URL发出一个POST请求。bodyType为POST数据的类型， body为POST数据，作为请求的主体
	如果参数body实现了io.Closer接口，它会在发送请求后被关闭。调用者有责任在读取完返回值resp的主体后关闭它。
 */
