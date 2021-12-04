package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "The target is wrong, for in attacking the tests, critics divert " +
		"attention from the fault that lies with ill informed or incompetent users."

	// 多条件匹配 -- 匹配ed和ing结尾的单词
	re := regexp.MustCompile(`\b[\w]+ing|\b[\w]+ed`)
	fmt.Println(re.FindAllString(str,-1))

	// 匹配200-1200范围内的数字
	re1 := regexp.MustCompile(`[2-9][0-9][0-9]|[1][0-1][0-9][0-9]|1200`)
	fmt.Println(re1.FindAllString("199",-1))
	fmt.Println(re1.FindAllString("564",-1))
	fmt.Println(re1.FindAllString("1200",-1))
	fmt.Println(re1.FindAllString("2250",-1))

}
