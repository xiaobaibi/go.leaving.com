package main

import "fmt"

/*
*
逻辑运算符
*/
func main() {
	fmt.Println(and())
	fmt.Println(or())
	fmt.Println(not())
}

func and() bool {
	return true && false
}

func or() bool {
	return true || false
}

func not() bool {
	return !true
}
