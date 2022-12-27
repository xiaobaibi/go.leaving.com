package main

import "fmt"

/*
*

	init函数是是初始化函数，当项目启动的时候会初始化启动，
	一个包中可以有多个init函数，同一个文件中有多个init函数时，会按上下顺序执行
*/
func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("main")
}
