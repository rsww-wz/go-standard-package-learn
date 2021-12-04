package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// 时间戳，单位秒
	t1 := now.Unix()

	// 纳秒时间戳
	//func (t Time) UnixNano() int64
	t2 := now.UnixNano()

	fmt.Printf("%T,%v\n",t1,t1)
	fmt.Printf("%T,%v\n",t2,t2)

	// 时间戳转时间对象
	//func Unix(sec int64, nsec int64) Time
	t3 := time.Unix(t1,0)   // 秒时间戳
	t4 := time.Unix(0,t2)   // 纳秒时间戳
	fmt.Printf("%T,%v\n",t3,t3)
	fmt.Printf("%T,%v\n",t4,t4)
}
