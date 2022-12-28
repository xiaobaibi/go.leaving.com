package main

import "fmt"

func main() {
	a()
	b()
}

func a() {
	fmt.Println("func a")
}
func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b is error  ")
		}
	}()
	panic("error")
}
