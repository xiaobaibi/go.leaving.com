package main

import (
	"fmt"
)

func main() {

	fmt.Println(funInt())
	fmt.Println(funBool())

}

func funInt() (int32, int32) {
	return 22, 32
}

func funBool() bool {
	return true
}
