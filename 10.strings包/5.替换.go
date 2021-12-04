package main

import (
	"fmt"
	"strings"
)

func main() {
	// 替换
	//func Replace(s, old, new string, n int) string
	//func ReplaceAll(s, old, new string) string
	text := "looking,taking,making,playing"
	res1 := strings.Replace(text,"ing","ed",1)
	res2 := strings.Replace(text,"ing","ed",3)
	res3 := strings.Replace(text,"ing","ed",-1)  // 全部替换
	res4 := strings.ReplaceAll(text,"ing","ed")  // 全部替换
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
}
