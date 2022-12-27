package main

import (
	"fmt"
	"strings"
)

var str = "abcadefag"

func main() {
	fmt.Println(str)
	// 长度
	fmt.Println(len(str))
	// 拼接
	fmt.Println(str + str)
	// 切割
	fmt.Println(strings.Split(str, "a"))
	// 判断前缀
	fmt.Println(strings.HasPrefix(str, "b"))
	// 判断后缀
	fmt.Println(strings.HasSuffix(str, "ag"))
	// 判断子字符串包含
	fmt.Println(strings.Contains(str, "ade"))
	// 判断子字符串位置
	fmt.Println(strings.Index(str, "ade"))
	// 判断子字符串位置
	fmt.Println(strings.LastIndex(str, "ade"))

}
