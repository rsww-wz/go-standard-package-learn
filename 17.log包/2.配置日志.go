package main

/*
	一般在包的init函数中配置log
	init函数会在main函数之前执行

所有日志都会记录日期和时间

配置函数也分为三类
	配置前缀
		func SetPrefix(prefix string)
		默认没有前缀

	配置输出
		func SetOutput(w io.Writer)
		一般是输出到文件中

	配置标志
		func SetFlags(flag int)
		默认就是3，也就是日期和时间

		const (
		Ldate         = 1 << iota     1 << 0 = 1
		Ltime                         1 << 0 = 2
		Lmicroseconds                 1 << 0 = 4
		Llongfile                     1 << 0 = 8
		Lshortfile                    1 << 0 = 16
		LUTC                          1 << 0 = 32
		Lmsgprefix                    1 << 0 = 64
		LstdFlags     = Ldate | Ltime  3
	)


 */

func main() {
	//log.SetPrefix()
	//log.SetOutput()
	//log.SetFlags()
}
