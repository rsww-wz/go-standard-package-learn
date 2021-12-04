/*
reader是一个接口，要实现它才能使用，有很多包都实现了这个接口
一般都是使用
Read方法用于读取数据
	参数：一个byte类型的切片
	返回值：读取到的字节大小

步骤
	打开文件
	读取数据
	关闭文件
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	address := `E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\1.txt`

	// 打开文件 -- 建立连接
	file,err := os.Open(address)
	if err != nil {
		fmt.Println(err)
		return
	}

	//闭关文件 -- 断开连接
	defer file.Close()

	//创建切片
	// 切片的大小一般是1024，也可以乘2或者乘4
	bs := make([]byte,16,16)

	/*
	// 第一次读取
	n,err := file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))

	// 第二次读取 -- 会接着上一次的位置开始读
	n,err = file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))

	// 第三次读取 -- 如果以及读完了，并且还有剩余空间，会重新开始读取本次取出来的数据

	n,err = file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))

	// 如果在进行第四次读取，不会再进行读取，err会返回EOF表示文件以及读取完毕
	n,err = file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	 */

	//读取的常用操作
	n := -1
	for {
		n,err = file.Read(bs)
		if n == 0||err == io.EOF {
			fmt.Println("文件读到了文件的末尾")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
