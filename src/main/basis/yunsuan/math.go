package main

import "fmt"

/*
*
数学运算
*/
func main() {
	fmt.Println(add(1, 2))
	fmt.Println(chenfa(2, 3))
}

func add(a, b int32) int32 {
	return a + b
}

func chenfa(a, b int32) int32 {
	return a * b
}
