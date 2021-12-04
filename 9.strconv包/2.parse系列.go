/*
Parse类函数用于转换字符串为给定类型的值
	ParseBool()
	ParseFloat()
	ParseInt()
	ParseUint()

注意
	任何类型转字符串都很容易，但是字符串转化为指定类型是有比较严格的条件的，是有可能失败的
	所以都需要返回err
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//func ParseBool(str string) (value bool, err error)
	//它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%v\n",b,b)


	//func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//返回字符串表示的整数值，接受正负号
	//base指定进制（2到36）
	//bitSize指定结果必须能无溢出赋值的整数类型
	i, err := strconv.ParseInt("-2", 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%v\n",i,i)

	//func ParseUint(s string, base int, bitSize int) (n uint64, err error)
	//类似ParseInt但不接受正负号，用于无符号整
	u, err := strconv.ParseUint("2", 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%v\n",u,u)

	//func ParseFloat(s string, bitSize int) (f float64, err error)
	//bitSize指定了期望的接收类型
	//函数会返回最为接近s表示值的一个浮点数
	f, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%v\n",f,f)
}
