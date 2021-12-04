package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "hello world"
	re := regexp.MustCompile(`\w`)

	//func (re *Regexp) FindAll(b []byte, n int) [][]byte
	//Find返回保管正则表达式re在b中的所有不重叠的匹配结果的[][]byte切片。如果没有匹配到，会返回nil
	res1 := re.FindAll([]byte(s),-1)
	fmt.Println(res1)
	for _,v:= range res1 {
		fmt.Print(string(v))
	}
	fmt.Println()

	//func (re *Regexp) FindAllString(s string, n int) []string
	//Find返回保管正则表达式re在b中的所有不重叠的匹配结果的[]string切片。如果没有匹配到，会返回nil
	res2 := re.FindAllString(s,-1)
	fmt.Println(res2)

	//func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
	//Find返回保管正则表达式re在b中的所有不重叠的匹配结果的起止位置的切片。如果没有匹配到，会返回nil
	s1 := "Good morning"
	re1 := regexp.MustCompile(`o`)
	res3 := re1.FindAllIndex([]byte(s1),-1)
	fmt.Println(res3)

	//func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
	//Find返回保管正则表达式re在b中的所有不重叠的匹配结果的起止位置的切片。如果没有匹配到，会返回nil
	res4 := re1.FindAllStringIndex(s1,-1)
	fmt.Println(res4)
}
