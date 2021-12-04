package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	dataMap := map[string]interface{}{
		"Name":  "王钢蛋",
		"age":   20,
		"rmb":   123.45,
		"sex":   true,
		"Hobby": []string{"抽中华", "喝牛栏山", "烫花卷头"},
	}

	// 打开文件
	dstFile, _ := os.OpenFile(`E:\Document\0-code\Golang\src\Golang\Go标准库\12.json包\data.json`, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	// 新建编码器
	//func NewEncoder(w io.Writer) *Encoder
	encoder := json.NewEncoder(dstFile)

	// 对数据进行编码
	err := encoder.Encode(dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("编码成功")

	defer dstFile.Close()
}
