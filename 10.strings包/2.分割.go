package main

import (
	"fmt"
	"strings"
)

func main() {
	//func Fields(s string) []string
	//返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串
	//如果字符串全部是空白或者是空字符串的话，会返回空切片
	s1 := "  foo bar  baz   "
	r1 := strings.Fields(s1)
	fmt.Println(r1)


	//func Split(s, sep string) []string
	//用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。
	//如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串
	str1 := "hello|world"
	str2 := "hello world"
	str3 := "你好，世界,hello world"
	fmt.Println(strings.Split(str1,"|"))
	fmt.Println(strings.Split(str2," "))
	fmt.Println(strings.Split(str3,""))     // 切分成每一个unicode码值一个字符串,也就是以字符为单位

	//func SplitN(s, sep string, n int) []string
	//和split一样，参数n决定返回的切片的数目
	//n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
	//n == 0: 返回nil
	//n < 0 : 返回所有的子字符串组成的切片
	fmt.Println(strings.SplitN(str3,"",8))
	fmt.Println(strings.SplitN(str3,"",-1))


	//func SplitAfter(s, sep string) []string
	//用从s中出现的sep后面切断的方式进行分割，会分割到结尾
	fmt.Println(strings.SplitAfter("你好，中国，这是我的故乡","中国"))
	fmt.Println(len(strings.SplitAfter("你好，中国，这是我的故乡","中国")))

	//func SplitAfterN(s, sep string, n int) []string
	//用从s中出现的sep后面切断的方式进行分割，会分割到结尾,参数n决定返回的切片的数目
	fmt.Println(strings.SplitAfterN(str3,"",4))
}
