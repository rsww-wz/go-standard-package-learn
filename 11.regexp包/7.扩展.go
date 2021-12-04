/*
func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte

	Expand返回新生成的将template添加到dst后面的切片
	在添加时，Expand会将template中的变量替换为从src匹配的结果
	match应该是被FindSubmatchIndex返回的匹配结果起止位置索引
	通常就是匹配src，除非你要将匹配得到的位置用于另一个[]byte

	在template参数里，一个变量表示为格式如：
	$name或${name}的字符串，其中name是长度>0的字母、数字和下划线的序列
	一个单纯的数字字符名如$1会作为捕获分组的数字索引
	其他的名字对应(?P<name>...)语法产生的命名捕获分组的名字
	超出范围的数字索引、索引对应的分组未匹配到文本、正则表达式中未出现的分组名，都会被替换为空切片

	$name格式的变量名，name会尽可能取最长序列
	$1x等价于${1x}而非${1}x，$10等价于${10}而非${1}0
	因此$name适用在后跟空格/换行等字符的情况，${name}适用所有情况

	如果要在输出中插入一个字面值'$'，在template里可以使用$$
 */

package main

func main() {
	//func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte


	//func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte
	//ExpandString类似Expand，但template和src参数为字符串。它将替换结果添加到切片并返回切片，以便让调用代码控制内存申请
}
