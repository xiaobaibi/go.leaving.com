package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum()
	random()
}

func sum() {
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = i
	}
	var sum int
	for _, num := range array {
		sum += num
	}
	fmt.Printf("和= %d", sum)
}

func random() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Printf("第%d随机数是%d", i, rand.Intn(100))
	}
}
