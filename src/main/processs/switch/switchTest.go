package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		ranNum := rand.Intn(7)
		fmt.Println(ranNum)
		go selectBlock(ranNum)
	}
}

func selectBlock(key int) {
	// 阻塞协程
	select {}
}

func selectNotBlock(key int) {
	// 阻塞协程
	select {

	default:
		fmt.Println("not block")
	}
}
