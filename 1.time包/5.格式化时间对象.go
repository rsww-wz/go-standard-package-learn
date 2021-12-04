package main

import (
	"fmt"
	"time"
)

func main() {
	// 格式化时间
	// 把时间对象按模板格式化成字符串
	// 模板：1月2日，15时4分5秒，06年（记：123456） -- 24小时制
	// 如果要12小时制需要指定PM
	now := time.Now()

	//func (t Time) Format(layout string) string
	//Format根据指定的格式返回t代表的时间点的格式化文本表示

	// 24小时制
	t1 := now.Format("2006-1-2 15:4")
	fmt.Printf("%T,%v\n",t1,t1)

	// 12小时制
	t2 := now.Format("2006/1/2 15:04:05 PM")
	fmt.Println(t2)

	t3 := now.Format("2006年1月2日-15:04:06")
	fmt.Println(t3)

	//func (t Time) String() string
	//String返回采用如下格式字符串的格式化时间
	//2006-01-02 15:04:05.999999999 -0700 MST
	s := now.String()
	fmt.Printf("%T,%v\n",s,s)
}
