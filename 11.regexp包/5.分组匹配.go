package main

import (
	"fmt"
	"regexp"
)

func main() {
	//func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
	//Find返回一个保管正则表达式re在b中的所有不重叠的匹配结果及其对应的（可能有的）分组匹配的结果的[][][]byte切片。
	//如果没有匹配到，会返回nil

	// 用的不多，和FindAllStringSubmatch用法一样

	//func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
	//Find返回一个保管正则表达式re在b中的所有不重叠的匹配结果及其对应的（可能有的）分组匹配的结果的[][]string切片。
	//如果没有匹配到，会返回nil。

	// 也就是不仅会匹配整个正则，还会匹配里面的分组，每匹配到一次正则，就作为一组
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))        //[["ab" ""]]
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))      //[["axxb" "xx"]]
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))    //[["ab" ""] ["axb" "x"]]
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))   //[["axxb" "xx"] ["ab" ""]]

	fmt.Println("-------------------------------------------------------")

	//func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
	//func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int

	// 按照 FindAllStringSubmatch匹配的结果把结果转化为索引
	re1 := regexp.MustCompile("a(x*)b")
	fmt.Println(re1.FindAllStringSubmatchIndex("-ab-", -1))
	fmt.Println(re1.FindAllStringSubmatchIndex("-axxb-", -1))
	fmt.Println(re1.FindAllStringSubmatchIndex("-ab-axb-", -1))
	fmt.Println(re1.FindAllStringSubmatchIndex("-axxb-ab-", -1))
	fmt.Println(re1.FindAllStringSubmatchIndex("-foo-", -1))

	//func (re *Regexp) FindSubmatch(b []byte) [][]byte
	//Find返回保管正则表达式re在b中的最左侧的一个匹配结果的起止位置的切片（显然len(loc)==2）。
	//匹配结果可以在输入流r的字节偏移量loc[0]到loc[1]-1（包括二者）位置找到。如果没有匹配到，会返回nil


	//func (re *Regexp) FindStringSubmatch(s string) []string
	//Find返回一个保管正则表达式re在b中的最左侧的一个匹配结果以及（可能有的）分组匹配的结果的[]string切片。
	//如果没有匹配到，会返回nil
	re2 := regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("%q\n", re2.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re2.FindStringSubmatch("-abzc-"))
}
