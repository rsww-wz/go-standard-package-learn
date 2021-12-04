package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// 读取数据
	srcFile, _ := os.Open(`E:\Document\0-code\Golang\src\Golang\Go标准库\12.json包\data.json`)

	// 新建解码器
	//func NewDecoder(r io.Reader) *Decoder
	decoder := json.NewDecoder(srcFile)

	dataMap := make(map[string]interface{})

	// 对数据进行解码
	err := decoder.Decode(&dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("解码成功")
	fmt.Println(dataMap)

	defer srcFile.Close()
}
