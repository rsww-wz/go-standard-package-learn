/*
功能
	如果要传的文件，比较大，是否有方法可以缩短耗时
	如果在文件传递过程中，程序被迫中断了，那么下次重启时，文件是否还需要重头开始
	传递文件的时候，支持暂停和恢复吗，即使这两个操作分布在程序进程被杀前后

实现
	Go语言通过Seek方法实现
思路
	记住上一次传递了多少数据，那我们可以创建一个临时文件，记录以及传递的数据量
	当恢复传递的时候，先从临时文件中读取上次已经穿钉的数据量
	然后通过Seek方法，设置到该读和该写的位置，再继续传递数据
 */
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	srcFile := `F:\图片\来源手机图片\花瓣网\河边1.jpg`
	destFile := srcFile[strings.LastIndex(srcFile,`\`)+1:]
	fmt.Println(destFile)
	tempFile := destFile+"temp.txt"
	fmt.Println(tempFile)

	file1,err := os.Open(srcFile)
	HandleErr(err)
	file2,err := os.OpenFile(destFile,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	HandleErr(err)
	file3,err := os.OpenFile(tempFile,os.O_CREATE|os.O_RDWR,os.ModePerm)
	HandleErr(err)

	defer file1.Close()
	defer file2.Close()

	// 步骤1：先读取临时文件中的数据在，seek
	file3.Seek(0,io.SeekStart)
	bs := make([]byte,100,100)
	n1,err := file3.Read(bs)
	//HandleErr(err)
	countStr := string(bs[:n1])
	count,err := strconv.ParseInt(countStr,10,64)
	//HandleErr(err)
	fmt.Println(count)

	//步骤2：设置读，写的位置
	file1.Seek(count,io.SeekStart)
	file2.Seek(count,io.SeekStart)
	data := make([]byte,1024,1024)
	n2 := -1               // 读取的数据量
	n3 := -1               // 写出的数据量
	total := int(count)    // 读取的总量

	// 步骤3：复制文件
	for{
		n2,err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制完毕")
			file3.Close()
			os.Remove(tempFile)
			break
		}
		n3,err = file2.Write(data[:n2])
		total += n3

		// 将复制的总量，存储到临时文件中
		file3.Seek(0,io.SeekStart)
		file3.WriteString(strconv.Itoa(total))
	}
}

func HandleErr(err error) {
	if err != nil{
		log.Fatal(err)
	}
}
