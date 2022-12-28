package main

import "fmt"

var array1 [5]int = [5]int{1, 2, 3, 4, 5}
var array2 = [5]int{122, 22, 33, 11, 2}
var array3 = [...]int{8, 8, 9, 00, 2}

func main() {
	test1()
}

func test1() {
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}
	fmt.Println("------------------")
	for index, num := range array1 {
		fmt.Println(index)
		fmt.Println(num)
	}
}
