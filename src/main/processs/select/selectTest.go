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
		switchT(ranNum)
	}

}

func switchT(key int) {
	switch key {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("Other")
	}
}
