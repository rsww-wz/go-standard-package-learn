package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// 在当前时间上加上多少时间
	t1 := now.Add(time.Hour * 24)
	t2 := now.Add(time.Hour * 24 * 30)
	t3 := now.Add(time.Hour * 24 * 365)

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)

	// 求两个时间之差
	t4 := now.Add(time.Millisecond)
	t5 := t4.Sub(now)
	fmt.Println(t5)

	// 判断时间前后
	// before：如果在now之前返回true，之后返回false
	// after:如果在now之前返回false,之后返回true
	t6 := now.Before(t4)
	t7 := now.After(t4)
	fmt.Println(t6)
	fmt.Println(t7)

	timeSub()
}


// 计算时间差
func timeSub() {
	// 计算时间差
	t1 := time.Now()
	t3 := t1.UnixNano()

	time.Sleep(1 * time.Second)
	time.Sleep(100 * time.Millisecond)

	t2 := time.Now()
	t4 := t2.UnixNano()

	t5 := t2.Sub(t1)
	t6 := t4 - t3

	fmt.Println("t1,t2的时间差是：",t5)
	// 纳秒和秒的数量级是：10e9
	fmt.Println("t1,t2的时间差是：",float64(t6)/10e8)
}