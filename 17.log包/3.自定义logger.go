package main

/*
默认的log只能根据记录信息，不能分级
自定义的logger就分级记录日志信息

一般自定义logger
	Trace    *log.Logger   // 几乎任何信息
	Info     *log.Logger   // 重要信息
	Warning  *log.Logger   // 警告
	Error    *log.Logger   // 错误

创建自定义logger
log.New()
	func New(out io.Writer, prefix string, flag int) *Logger

返回*logger对象可以调用记录日志的函数，入print，fatal，panic等

常用做法

Trace = log.New(ioutil.Discard,
		TRACE: ",
		log.Ldata|log.Ltime|log.Lshortfile)

	说明:type discard struct{}              // io包中
    	var Discard Writer = discard{}     // io包中
		var Discard io.Writer = io.Discard // ioutil包中


Info = log.New(os.Stdout,
		INFO: ",
		log.Ldata|log.Ltime|log.Lshortfile)

Warning = log.New(os.Stdout,
		WARNING: ",
		log.Ldata|log.Ltime|log.Lshortfile)

Error = log.New(io.MultiWriter(file,os.Stderr),
		ERROR: ",
		log.Ldata|log.Ltime|log.Lshortfile)

*/

func main() {
	//log.Logger{}
	//log.New()
	//var Discard io.Writer = devNull(0)
}
