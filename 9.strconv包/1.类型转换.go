/*
Atoi -- 字符串转整型
itoa -- 整型转字符串

记忆方法
	a：代表字符串，因为C语言中没有字符串，使用字符数组表示的，也就是array，所以字符串用a表示
	i：代表整型
	to：转换

注意
	atoi有err返回值，因为字符串转整型有可能失败
	itoa没有err返回值，因为整型一定可以转换为字符串
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi() -- string -> int
	//func Atoi(s string) (i int, err error)

	s1 := "100"
	value, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", value, value) //type:int value:100
	}

	// itoa() -- int -> string
	//func Itoa(i int) string
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"

}
