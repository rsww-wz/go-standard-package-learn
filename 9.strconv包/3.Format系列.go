/*
Format系列函数实现了将给定类型数据格式化为string类型数据的功能
也就是给定类型转换为字符串类型，都没有返回err
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//func FormatBool(b bool) string
	//根据b的值返回”true”或”false“
	f := strconv.FormatBool(true)
	fmt.Println(f)

	//func FormatInt(i int64, base int) string
	//返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字
	i1 := strconv.FormatInt(100,2)
	i2 := strconv.FormatInt(100,8)
	i3 := strconv.FormatInt(100,10)
	i4 := strconv.FormatInt(100,16)
	fmt.Println(i1,i2,i3,i4)

	//func FormatUint(i uint64, base int) string
	//是FormatInt的无符号整数版本
	u1 := strconv.FormatUint(346,10)
	u2 := strconv.FormatUint(346,16)
	fmt.Println(u1,u2)

	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	// fmt表示格式,prec控制精度（排除指数部分）,bitSize表示f的来源类型
	f1 := strconv.FormatFloat(3.14,'E',-1,64)
	fmt.Println(f1)
}
