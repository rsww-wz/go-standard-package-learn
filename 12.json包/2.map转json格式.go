/*
Marshal(编码)
	把Go struct转化为json格式

Marshalldent(编码)
	带缩进

Unmarshal(解码)
	把json转化为Go struct

总结
	针对string或者bytes
		marshal   -> string
		unmarshal -> string

	针对stream
		Encode -> stream,把数据写入到io.Writer
		Decode -> stream,从io.Reader读取数据
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 定义字典数据
	studentInfo := map[string]interface{} {
		"name":"rs",
		"age": 21,
		"gender": true,
		"ID": "6635",
		"address":nil,
	}

	// 把字典转为JSON格式
	data,err := json.Marshal(studentInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
