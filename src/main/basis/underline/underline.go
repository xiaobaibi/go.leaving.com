package main

import "fmt"
import _ "go.leaving.com/src/main/initp"

/*
*
1 ： 下滑线在import中 ：形如 import _ "go.leaving.com/src/main/initp" 不使用其中的函数，可以触发执行其包下的init方法
2 ： 在函数调用返回值中可以利用下滑线做忽略【忽略变量】
*/
func main() {

	fmt.Println("underline")
	a, _ := underLine(1, 2)
	fmt.Println(a)

}

func underLine(a, b int32) (int32, int32) {
	return a, b
}
