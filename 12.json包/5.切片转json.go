package main

import (
	"encoding/json"
	"fmt"
)

// 切片放struct，切片放map都可以
func main() {
	dataMap1 := map[string]interface{}{
		"Name":  "王钢蛋",
		"Hobby": []string{"抽中华", "喝牛栏山", "烫花卷头"},
	}

	dataMap2 := map[string]interface{}{
		"Name":  "王铁蛋",
		"Hobby": []string{"抽玉溪", "喝五粮液", "烫杀马特"},
	}

	dataMap3 := map[string]interface{}{
		"Name":  "王铜蛋",
		"Hobby": []string{"抽小熊猫", "喝剑南春", "烫鸡冠头"},
	}

	dataSlice := make([]map[string]interface{}, 0)
	dataSlice = append(dataSlice, dataMap1, dataMap2, dataMap3)

	data, err := json.Marshal(dataSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
