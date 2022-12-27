package main

import "fmt"

/*
*

	在方法上面进行传参的时候 如果参数以*修饰，则在传入的时候需要传入变量的批量位置。以&修饰
	如果传入指针后。变量被修改，则该参数会被永久修改
*/
func main() {
	status := [5]int{1, 2, 3, 4, 5}
	fmt.Println("main : ", status)
	changeStatus(&status)
	fmt.Println("main : ", status)
	changeStatus2(status)
	fmt.Println("main : ", status)
}

func changeStatus(status *[5]int) {
	fmt.Println("changeStatus : ", status)
	status[0] = 55
	fmt.Println("changeStatus : ", status)
}

func changeStatus2(status [5]int) {
	fmt.Println("changeStatus2 : ", status)
	status[0] = 66
	fmt.Println("changeStatus2 : ", status)
}
