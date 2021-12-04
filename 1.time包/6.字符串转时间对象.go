package main

import (
	"fmt"
	"time"
)

func main() {
	// 把字符串转换成时间对象
	t1 := "1999年10月1日"
	// 模板要和时间的格式一样才能转化成时间对象
	t2,err := time.Parse("2006年1月2日",t1)
	if err != nil {
		fmt.Println("err",err)
		return
	}
	fmt.Printf("%T,%v\n",t2,t2)

	// 获取时间的字符串
	now := time.Now()

	t3 := now.String()
	fmt.Printf("%T,%v\n",t3,t3)
}
