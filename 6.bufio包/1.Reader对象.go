/*
功能
	通过缓冲来提高效率
原理
	io操作本身的效率并不低，低的是频繁的访问本地磁盘的文件
	所以bufio就是提供了缓冲区，读和写都是先在缓冲区中
	最后再读写文件，来降低访问本地磁盘的次数，从而提高效率

理解
	简单来说就是，把文件读取进内存之后再读取的时候就可以避免文件系统io从而提高速度
	同理，在进行写操作票的时候，把文件写入内存，然后又缓冲区写入文件系统

	缓冲区是为了避免多次访问磁盘

bufio.Reader是bufio对io.Reader的封装
思路
	当缓存区有内容时，将缓存区内容全部填入p并清空缓存区
	当缓存区没有内容时并且len(p)>len(buf)，要读取的内容比缓存区还要大，直接去文件读取即可(缓存区没有意义了)
	当缓存区没有内容时并且len(p)<len(buf)，要读取的内容比缓存区小，缓存区从文件读取内容充满缓存区并jlp填满
	以后再次读取时，缓存区有内容，将缓存区内容全部填入p并清空缓存区(此时和情况1一样)

 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	address := `E:\0-code\Go\src\Golang\Go标准库\3.文件操作\test\1.txt`
	file,err := os.Open(address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//func NewReaderSize(rd io.Reader, size int) *Reader
	//NewReaderSize创建一个具有最少有size尺寸的缓冲、从r读取的*Reader
	//如果参数r已经是一个具有足够大缓冲的* Reader类型值，会返回r

	//func NewReader(rd io.Reader) *Reader
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader    默认是4096byte = 4K
	b1 := bufio.NewReader(file)
	p := make([]byte,1024)

	//func (b *Reader) Read(p []byte) (n int, err error)
	//本方法返回写入p的字节数。本方法一次调用最多会调用下层Reader接口一次Read方法，因此返回值n可能小于len(p)
	//读取到达结尾时，返回值n将为0而err将为io.EOF
	n1,err := b1.Read(p)
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))

	//func (b *Reader) Buffered() int
	//Buffered返回缓冲中现有的可读取的字节数
	fmt.Println(b1.Buffered())
}
