package main

import (
	"fmt"
	"go.leaving.com/src/main/basis/structs"
)

func main() {
	normalStruct()
	nimingStruct()
	newStruct()
	initStruct()
	structureStruct()
}

/*
* 普通结构体
 */
func normalStruct() {
	fmt.Println("----------------------普通结构体--------------------------")
	var person structs.Person
	person.Name = "jianghai"
	person.Age = 30
	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Printf("%v", person)
}

/*
*
匿名结构体
*/
func nimingStruct() {
	fmt.Println("----------------------匿名结构体--------------------------")
	var user struct {
		Name string
		Age  int
	}
	user.Name = "Alice"
	user.Age = 22
	fmt.Printf("%v", user)
}

func newStruct() {
	fmt.Println("----------------------New结构体--------------------------")
	var p = new(structs.Person)
	p.Name = "River"
	p.Age = 22
	// 可以看见打印出来的是地址
	fmt.Printf("%T\n", p)
	fmt.Printf("p2=%#v\n", p)

	fmt.Println("----------------------&结构体--------------------------")
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	var p2 = &structs.Person{}
	p2.Name = "Felix"
	p2.Age = 18
	// 可以看见打印出来的是地址
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)
}

/*
*
结构体初始化
*/
func initStruct() {
	fmt.Println("----------------------初始化--------------------------")

	var p1 structs.Person
	var p2 = new(structs.Person)
	var p3 = &structs.Person{}
	p4 := structs.Person{}

	fmt.Printf("p1=%#v\n", p1)
	fmt.Printf("p2=%#v\n", p2)
	fmt.Printf("p3=%#v\n", p3)
	fmt.Printf("p4=%#v\n", p4)
}

/*
*
构造函数
*/
func structureStruct() {
	fmt.Println("----------------------构造函数--------------------------")
	// 指针类型
	person := structs.NewPerson("jianghai", 33)
	fmt.Printf("person=%#v\n", person)
	// 非指针类型
	person2 := structs.NewPerson2("Alice")
	fmt.Printf("person2=%#v\n", person2)
	person.PrintName()

	fmt.Println(person.Age)
	person.SetAge(10)
	fmt.Println(person.Age)
	person.SetAge2(11)
	fmt.Println(person.Age)

}
