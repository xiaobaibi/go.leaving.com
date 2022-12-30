package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := funcA(5); err != nil {
		fmt.Println("出错了")
	}
	if err := funcA(-1); err != nil {
		fmt.Println("出错了")
	}
}

func funcA(a int) error {
	if a < 0 {
		return errors.New("参数小于0")
	}
	fmt.Println(a)
	return nil
}
