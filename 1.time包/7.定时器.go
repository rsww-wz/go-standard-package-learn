package main

import (
	"fmt"
	"time"
)

func main() {
	//定时器，它的本质是一个通道
	// 定义一个1秒间隔的定时器
	ticker := time.Tick(time.Second)
	for i := range ticker {
		// 每秒都会执行任务
		fmt.Println(i)
	}
}
