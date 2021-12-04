package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间，返回的是一个对象
	now := time.Now()
	fmt.Println(now)

	// 获取指定时间对象
	// 参数：年月日，时分秒，毫秒，时区
	t1 := time.Date(2008,1,1,15,30,45,0,time.Local)   // date函数
	fmt.Println(t1)

	// 获取日期，date方法
	years,months,days := now.Date()
	fmt.Println(years,months,days)

	// 获取时间，clock方法
	hours,minutes,seconds := now.Clock()
	fmt.Println(hours,minutes,seconds)

	// 获取时间对象指定部分
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%T,%v\n",now,now)
	fmt.Printf("%T,%v\n",day,day)
	fmt.Println(year,month,day,hour,minute,second)

	// 获取距离今年过去了多少天
	yearDay := now.YearDay()
	fmt.Printf("2021年已经过去了%v天\n",yearDay)
}
