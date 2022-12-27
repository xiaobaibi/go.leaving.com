package main

import "fmt"

/*
*
关系运算
*/
func main() {
	fmt.Println(equ("a", "a"))
}

func equ(str1, str2 string) bool {
	return str1 == str2
}
