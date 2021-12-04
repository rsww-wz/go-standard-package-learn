package main

/*
type Client struct
    Transport RoundTripper
		Transport指定执行独立、单次HTTP请求的机制
		如果Transport为nil，则使用DefaultTransport

		Client实例的Transport字段必须支持CancelRequest方法，
    	否则Client会在试图用Head、Get、Post或Do方法执行请求时返回错误。
    	本类型的Transport字段默认值（DefaultTransport）支持CancelRequest方法

	CheckRedirect func(req *Request, via []*Request) error
		CheckRedirect指定处理重定向的策略
		如果CheckRedirect不为nil，客户端会在执行重定向之前调用本函数字段。
		参数req和via是将要执行的请求和已经执行的请求（切片，越新的请求越靠后）
		如果CheckRedirect返回一个错误，本类型的Get方法不会发送请求req
		而是返回之前得到的最后一个回复和该错误。（包装进url.Error类型里）
		如果CheckRedirect为nil，会采用默认策略：连续10此请求后停止

    Jar CookieJar
		Jar指定cookie管理器
		如果Jar为nil，请求中不会发送cookie，回复中的cookie会被忽略

    Timeout time.Duration
		Timeout为零值表示不设置超时
		Timeout指定本类型的值执行请求的时间限制
		该超时限制包括连接时间、重定向和读取回复主体的时间
		计时器会在Head、Get、Post或Do方法返回后继续运作并在超时后中断回复主体的读取

Client类型代表HTTP客户端。它的零值（DefaultClient）是一个可用的使用DefaultTransport的客户端
Client的Transport字段一般会含有内部状态（缓存TCP连接），因此Client类型值应尽量被重用而不是每次需要都创建新的
Client类型值可以安全的被多个go程同时使用

Client类型的层次比RoundTripper接口（如Transport）高，还会管理HTTP的cookie和重定向等细节

 */
func main() {

}
