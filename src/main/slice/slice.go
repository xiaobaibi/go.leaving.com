package main

import "fmt"

/*
*
 */
func main() {
	initSlice()
	makeSlice()
	copySlice()
	getMemory()
}

/*
*
声明
*/
func initSlice() {
	// 声明方式1 ：
	var arr1 []int
	if arr1 == nil {
		fmt.Println("空切片")
	} else {
		fmt.Println("非空切片")
	}

	// 声明方式2 ：
	arr2 := []int{}
	fmt.Println(arr2)
	// 声明方式3 ：
	var arr3 []int = make([]int, 0)
	fmt.Println(arr3)
	// 声明方式4 ：
	var arr4 []int = make([]int, 0, 0)
	fmt.Println(arr4)
	// 声明方式5 ：
	arr5 := [5]int{1, 2, 3, 4, 5}

	// 截取【包前不包后】
	fmt.Println(arr5[1:2])

}

func makeSlice() {
	//10:切片的长度
	//11：切片的容量
	arr1 := make([]int, 10, 20)
	for i := 0; i < 10; i++ {
		arr1[i] = i
	}
	for index, sum := range arr1 {
		fmt.Printf("角标：%d，值%d", index, sum)
	}
	for i := 10; i < 15; i++ {
		arr1 = append(arr1, i)
	}
	fmt.Println()
	for index, sum := range arr1 {
		fmt.Printf("角标：%d，值%d", index, sum)
	}
	fmt.Println()
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))
}

func copySlice() {
	arr := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	arr2 := make([]int, 10)
	copy(arr2, arr)
	for _, num := range arr2 {
		fmt.Println(num)
	}
}

func getMemory() {
	var a *int
	*a = 1
	fmt.Println(a)

	//var b map[string]int
	//b["a"]=1
	//fmt.Println(b)
}
