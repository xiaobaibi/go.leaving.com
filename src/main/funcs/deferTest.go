package main

import "fmt"

/*
*
def 相当于Java里面的Finally
用于:
 1. 关闭文件句柄
 2. 锁资源释放
 3. 数据库连接释放
*/
func main() {
	defer fmt.Println("abc")

	fmt.Println("e")
	fmt.Println("f")
}
