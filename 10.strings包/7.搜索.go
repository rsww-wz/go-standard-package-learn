package main

import (
	"fmt"
	"strings"
)

func main() {
	//func Index(s, sep string) int
	//子串sep在字符串s中第一次出现的位置，不存在则返回-1
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	//func LastIndex(s, sep string) int
	//子串sep在字符串s中最后一次出现的位置，不存在则返回-1
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))

	//func Repeat(s string, count int) string
	//返回count个s串联的字符串
	s1 := "s"
	r1 := strings.Repeat(s1,10)
	fmt.Println(r1)
}
