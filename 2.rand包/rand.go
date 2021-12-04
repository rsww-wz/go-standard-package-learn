package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置种子，需要设置一次
	// 如果种子参数一样，每次运行程序产生的随机数都是一样的
	//rand.Seed(0)
	// 当时间戳作为种子，每次产生的随机数就不一样了
	rand.Seed(time.Now().UnixNano())

	// 生成int64范围内的随机非负整数
	fmt.Println(rand.Int())

	// 生成指定范围内的随机整数
	s1 := rand.Intn(100)           // [0-100)
	s2 := rand.Intn(28) + 18       // [28-46)
	fmt.Println(s1)
	fmt.Println(s2)

	// 生成0-1内的随机浮点数
	fmt.Println(rand.Float64())  // [0-1)

	// 生成指定范围内的随机浮点数
	// 乘法扩大范围，加法抬高下限
	r1 := rand.Float64()*100       // [0-100)
	r2 := rand.Float64()+5         // [5-6)
	r3 := (rand.Float64()*100)+18  // [18-118)
	fmt.Println("0-100内随机数：",r1)
	fmt.Println("5-6内随机数：",r2)
	fmt.Println("18-118内随机数：",r3)

	// 生成随机数组
	var lst []int
	t1 := time.Now()
	t3 := t1.UnixNano()

	// 有的数字会重复
	for i:=0;i<100000000;i++ {
		lst = append(lst,rand.Intn(100000000))
	}

	t2 := time.Now()
	t4 := t2.UnixNano()
	fmt.Printf("生成%d个随机数的时间戳是：%d\n",len(lst),t4-t3)
	fmt.Printf("生成%d个随机数的耗时是：%v\n",len(lst),t2.Sub(t1))
}
