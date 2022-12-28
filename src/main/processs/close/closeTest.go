package main

import (
	"fmt"
	"strings"
)

// 闭包 = 函数 + 外层变量的引用
func main() {
	func1 := makSuffix(".txt")
	ret := func1("斗破")
	fmt.Println(ret)
}

// 闭包函数
func makSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
