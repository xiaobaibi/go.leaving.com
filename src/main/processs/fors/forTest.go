package main

import "fmt"

func main() {
	fori()
	//for2()
	//forWhile()
	//forTrue()

}

func fori() {
	str := "abc"
	for a, b := 0, len(str); a < b; a++ {
		fmt.Println(a)
	}
}

func for2() {
	str := "abc"
	for len(str) > 2 {
		fmt.Println(str)
	}
}

func forWhile() {
	str := "abc"
	for {
		fmt.Println(str)
	}
}

func forTrue() {
	str := "abc"
	for true {
		fmt.Println(str)
	}
}
