package main

func main() {

}

/*
func NewRequest(method, urlStr string, body io.Reader) (*Request, error)
	NewRequest使用指定的方法、网址和可选的主题创建并返回一个新的*Request
	如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body
	并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭

func ReadRequest(b *bufio.Reader) (req *Request, err error)
	ReadRequest从b读取并解析出一个HTTP请求。（本函数主要用在服务端从下层获取请求）

func (r *Request) UserAgent() string
	UserAgent返回请求中的客户端用户代理信息（请求的User-Agent头）

func (r *Request) Referer() string
	Referer返回请求中的访问来路信息

func (r *Request) AddCookie(c *Cookie)
	AddCookie向请求中添加一个cookie

func (r *Request) Write(w io.Writer) error
	Write方法以有线格式将HTTP/1.1请求写入w（用于将请求写入下层TCPConn等）

func (r *Request) WriteProxy(w io.Writer) error
	WriteProxy类似Write但会将请求以HTTP代理期望的格式发送。

func (r *Request) Cookies() []*Cookie
	Cookie返回请求中名为name的cookie，如果未找到该cookie会返回nil, ErrNoCookie

func (r *Request) ParseForm() error
	ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段

func (r *Request) FormValue(key string) string
	FormValue返回key为键查询r.Form字段得到结果[]string切片的第一个值

func (r *Request) PostFormValue(key string) string
	PostFormValue返回key为键查询r.PostForm字段得到结果[]string切片的第一个值
	如果必要，本函数会隐式调用ParseMultipartForm和ParseForm

func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	FormFile返回以key为键查询r.MultipartForm字段得到结果中的第一个文件和它的信息

func (r *Request) MultipartReader() (*multipart.Reader, error)
	如果请求是multipart/form-data POST请求，MultipartReader返回一个multipart.Reader接口，否则返回nil和一个错误
	使用本函数代替ParseMultipartForm，可以将r.Body作为流处理
 */
