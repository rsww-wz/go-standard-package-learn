/*
读取json
	需要一个解码器
	dec := json.NewDecoder(r.Body)
		参数需要实现reader接口

	在解码器上进行解码
	dec := Decode(&query)

如果结构体需要导出，字段名首字母通常是需要大写的

此时可以在结构体后面用反引号：`json:"键名"`，表示导出的字段名

写入json
	需要一个编码器
	enc := json.NewEncoder(w)
		参数需要实现writer接口

	编码：enc.Encode(results)

*/

/*
	对json数据进行解码
	func Unmarshal(data []byte, v interface{}) error
		第一个参数：要转化的字节切片
		第二个参数：转化后数据的接收对象,是一个空接口类型
 */

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 字符串格式的json数据
	jsonStr := `{"gender":true,"num":123,"str":"Hello"}`

	// 把json字符串数据转化成字节类型
	jsonByte := []byte(jsonStr)

	value1 := transMap(jsonByte)
	fmt.Println(value1)

	value2 := transStruct(jsonByte)
	fmt.Println(value2)

}

func transMap (bytes []byte) map[string]interface{}{
	var mapData map[string]interface{}
	err := json.Unmarshal(bytes,&mapData)
	if err != nil{
		panic(err)
	}
	return mapData
}

func transStruct (bytes []byte) data {
	var structData data
	err := json.Unmarshal(bytes,&structData)
	if err != nil{
		panic(err)
	}
	return structData
}

// 结构体成员首字母需要大写才能读取
type data struct {
	Gender bool   `json:"gender"`
	Num int       `json:"num"`
	Str string    `json:"str"`
}
