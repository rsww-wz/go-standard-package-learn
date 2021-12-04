package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "hello world"
	re := regexp.MustCompile(`\w`)
	re1 := regexp.MustCompile(`\d`)

	//func (re *Regexp) Find(b []byte) []byte
	//Find返回保管正则表达式re在b中的最左侧的一个匹配结果的[]byte切片。如果没有匹配到，会返回nil
	res1 := re.Find([]byte(s))
	res2 := re1.Find([]byte(s))
	fmt.Println(string(res1))
	fmt.Println(string(res2))  // nil

	//func (re *Regexp) FindString(s string) string
	//Find返回保管正则表达式re在b中的最左侧的一个匹配结果的字符串。如果没有匹配到，会返回""
	res3 := re.FindString(s)
	res4 := re1.FindString(s)
	fmt.Println(res3)
	fmt.Println(res4)          //""
	fmt.Println("-------------------------------------------------------")

	//func (re *Regexp) FindIndex(b []byte) (loc []int)
	//Find返回保管正则表达式re在b中的最左侧的一个匹配结果的起止位置的切片（显然len(loc)==2）。
	//匹配结果可以通过起止位置对b做切片操作得到：b[loc[0]:loc[1]]。如果没有匹配到，会返回nil。
	re2 := regexp.MustCompile(`ll`)
	loc := re2.FindIndex([]byte(s))
	fmt.Println(loc)

	//func (re *Regexp) FindStringIndex(s string) (loc []int)
	//Find返回保管正则表达式re在b中的最左侧的一个匹配结果的起止位置的切片（显然len(loc)==2）。
	//匹配结果可以通过起止位置对b做切片操作得到：b[loc[0]:loc[1]]。如果没有匹配到，会返回nil。
	loc1 := re2.FindStringIndex(s)
	fmt.Println(loc1)
}
