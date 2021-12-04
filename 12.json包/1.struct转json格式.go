/*
json 是一种比xml更轻量级的数据交互格式，易于读和编写，也易于程序的解析和生成
json表现形式一般为键值对，这种json传输称为较为理想的跨平台的数据交互语言

开发者可以用json传输简单的字符串，数字，布尔值，也可以传输一个数组，或者一个更复杂的结构
在web开发中json被广泛的应用于web服务器和程序之间的数据通信

Go语言内置了标准库 encoding/json，开发者可以轻松使用语言程序生成json格式的数据
Go语言中的结构体和map都可以转换为json格式

json格式：value需要用双引号
*/

// 数据类型映射
/*
		Go		json
布尔值   bool	boolean
数值		float64	数值
浮点型	string	strings
空		nil		null

未知结构的json
map[string]interface{} :可以存储任意json对象
[]interface{}		   :可以存储任意json数组
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 结构体
	type Company struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country string `json:country`
	}

	// 实例化
	LG := Company{
		ID:6635,
		Name: "rw",
		Country: "韩国",
	}


	/*
	func Marshal(v interface{}) ([]byte, error)
		参数是一个空接口，可以是任何类型
		返回的类型是bytes类型
	 */

	LGdata,err := json.Marshal(LG)
	if err != nil {
		panic(err)
	}

	fmt.Printf("LG_json的数据类型是：%T\n",LGdata)

	// 转换成字符串类型
	data := string(LGdata)
	fmt.Printf("%T\n",data)
	fmt.Println(data)
}


