package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 向ch1中添加数据
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// ch1读取数据
	go func() {
		for {
			num, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- num * 2
		}
		close(ch2)
	}()
	// 循环读取数据
	for numDouble := range ch2 {
		fmt.Println(numDouble)
	}
}
