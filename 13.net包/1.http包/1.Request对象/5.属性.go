package main

func main() {

}

/*
type Request struct {
	Method string
		Method指定HTTP方法（GET、POST、PUT等）。对客户端，""代表GET。

	URL *url.URL
		URL在服务端表示被请求的URI，在客户端表示要访问的URL。
		在服务端，URL字段是解析请求行的URI（保存在RequestURI字段）得到的，对大多数请求来说，除了Path和RawQuery之外的字段都是空字符串
		在客户端，URL的Host字段指定了要连接的服务器，而Request的Host字段（可选地）指定要发送的HTTP请求的Host头的值

	Proto      string // "HTTP/1.0"
		接收到的请求的协议版本。本包生产的Request总是使用HTTP/1.1
		ProtoMajor int    // 1
		ProtoMinor int    // 0

	Header http.Header
		Header字段用来表示HTTP请求的头域。如果头域（多行键值对格式）为：
		HTTP规定头域的键名（头名）是大小写敏感的，请求的解析器通过规范化头域的键名来实现这点。
		在客户端的请求，可能会被自动添加或重写Header中的特定的头，参见Request.Write方法。

	Body io.ReadCloser
		Body是请求的主体
		在客户端，如果Body是nil表示该请求没有主体买入GET请求,Client的Transport字段会负责调用Body的Close方法
		在服务端，Body字段总是非nil的；但在没有主体时，读取Body会立刻返回EOF,Server会关闭请求的主体，ServeHTTP处理器不需要关闭Body字段

	ContentLength int64
		ContentLength记录相关内容的长度
		如果为-1，表示长度未知，如果>=0，表示可以从Body字段读取ContentLength字节数据
		在客户端，如果Body非nil而该字段为0，表示不知道Body的长度

	TransferEncoding []string
		TransferEncoding按从最外到最里的顺序列出传输编码，空切片表示"identity"编码
		本字段一般会被忽略。当发送或接受请求时，会自动添加或移除"chunked"传输编码

	Close bool
		Close在服务端指定是否在回复请求后关闭连接，在客户端指定是否在发送请求后关闭连接。

	Host string
		在服务端，Host指定URL会在其上寻找资源的主机,根据RFC 2616，该值可以是Host头的值，或者URL自身提供的主机名
		在客户端，请求的Host字段（可选地）用来重写请求的Host头,如过该字段为""，Request.Write方法会使用URL字段的Host

	Form url.Values
		Form是解析好的表单数据，包括URL字段的query参数和POST或PUT的表单数据
		本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代

	PostForm url.Values
		PostForm是解析好的POST或PUT的表单数据。
		本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。

	MultipartForm *multipart.Form
		MultipartForm是解析好的多部件表单，包括上传的文件
		本字段只有在调用ParseMultipartForm后才有效
		在客户端，会忽略请求中的本字段而使用Body替代


	RequestURI string
		RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
		一般应使用URI字段，在客户端设置请求的本字段会导致错误

	TLS *tls.ConnectionState
		LS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息
		本字段不是ReadRequest函数填写的
		对启用了TLS的连接，本包的HTTP服务器会在调用处理器之前设置TLS字段，否则将设TLS为nil
		客户端会忽略请求中的TLS字段

}
 */
