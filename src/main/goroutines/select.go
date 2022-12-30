package main

import (
	"fmt"
	"time"
)

func chA(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "chaA"
}

func chB(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "chaB"
}

/*
*
哪个Chan好了就执行哪个
*/
func main() {
	cha := make(chan string)
	chb := make(chan string)
	go chA(cha)
	go chB(chb)
	select {
	case chaResult := <-cha:
		fmt.Println("cha = ", chaResult)
	case chbResult := <-chb:
		fmt.Println("chb = ", chbResult)
	}
}
