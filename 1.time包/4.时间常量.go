package main

import (
	"fmt"
	"time"
)

func main() {
	// 时间常量，以纳秒为单位
	t1 := time.Hour
	t2 := time.Minute
	t3 := time.Second
	t4 := time.Millisecond
	t5 := time.Microsecond

	fmt.Printf("%T,%v\n",t1,t1)
	fmt.Printf("%T,%v\n",t2,t2)
	fmt.Printf("%T,%v\n",t3,t3)
	fmt.Printf("%T,%v\n",t4,t4)
	fmt.Printf("%T,%v\n",t5,t5)
}
