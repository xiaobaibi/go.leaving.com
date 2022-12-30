package main

import "fmt"

func receive(ch chan string) {
	str := <-ch
	fmt.Println(str)
}

func main() {
	ch := make(chan string)
	defer close(ch)
	go receive(ch)
	ch <- "jianghai"
	fmt.Println("发送成功")
}
