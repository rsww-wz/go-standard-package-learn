package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dirName := `F:\教程\python爬虫`
	listFiles(dirName,0)
}
func listFiles(dirName string,level int) {
	// level用来记录当前递归的层次，生成带有层次感的空格
	s := "|---"
	for i:=0;i<level;i++{
		s = "|  " + s
	}

	fileInfos,err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _,fi := range fileInfos {
		filename := dirName+`\`+fi.Name()
		fmt.Printf("%s%s\n",s,filename)
		if fi.IsDir() {
			// 递归调用方法
			listFiles(filename,level+1)
		}
	}
}
