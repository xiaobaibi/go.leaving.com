package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type Animal interface {
	Sayer
	Mover
}

type Dog struct {
}

func (dog Dog) say() {
	fmt.Println("汪")
}

func (dog Dog) move() {
	fmt.Println("狗爬")
}

type Cat struct {
}

func (cat Cat) say() {
	fmt.Println("瞄")
}
func main() {
	var sayer Sayer
	dog := Dog{}
	cat := Cat{}
	sayer = dog
	sayer.say()
	sayer = cat
	sayer.say()
	fmt.Println("------------")
	var animal Animal
	d := Dog{}
	animal = d
	animal.say()
	animal.move()
}
