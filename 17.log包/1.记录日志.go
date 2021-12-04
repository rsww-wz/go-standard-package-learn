package main

import "log"

/*
日志函数分为3类

第一类
	log.Print("")
	log.Printf("")
	log.Println("")
	只会打印内容
	记录一般日志

第二类
	log.Fatal("")
	log.Fataf("")
	log.Fatalln("")
	会先打印内容，然后调用os.exit(1)退出函数
	后面的代码不会执行
	记录重大错误


第三类
	log.Panic("")
	log.Panicf("")
	log.Panicln("")
	函数发生恐慌，先打印内容，然后再打印错误信息，最后终止程序
*/

func main() {
	//log.Println("1234")
	//log.Fatalln("1234")
	log.Fatalln("1234")
}
