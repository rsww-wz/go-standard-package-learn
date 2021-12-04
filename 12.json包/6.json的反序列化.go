package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"Name":"于谦","Age":50,"Rmb":123.45,"Sex":true,"Hobby":["抽烟","喝酒","烫头"]}`

	//func Unmarshal(data []byte, v interface{}) error
	dataMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dataMap)
}
