package main

import "fmt"

func main() {
	ch := make(chan int)
	// 循环10次往向Chan中添加数据,添加完关闭
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	// 循环读取数据
	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}
