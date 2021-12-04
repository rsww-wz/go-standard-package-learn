/*
	ReadDir可以读取一个目录的内容
	但是只能读取一个层目录，即字文件和字目录
 */
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dirName := `F:\电影\战争`
	fileInfos,err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(fileInfos))
	for i:=0;i<len(fileInfos);i++ {
		//fmt.Printf("%T\n",fileInfos[i])
		fmt.Printf("第%d个:名称：%s，是否是目录：%t\n",i,fileInfos[i].Name(),fileInfos[i].IsDir())
	}
}
