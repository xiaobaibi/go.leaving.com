package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go waitString(ch)
	for str := range ch {
		fmt.Println(str)
		time.Sleep(time.Second)
	}
}

func waitString(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("hello")
		default:
			fmt.Println("full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
