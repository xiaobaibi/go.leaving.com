package structs

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string `json:"addr"`
}

// NewPerson 构造函数/
func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func NewPerson2(name string) Person {
	return Person{Name: name}
}

// PrintName 打印名称/
func (p Person) PrintName() {
	fmt.Println(p.Name)
}

// 修改年龄
func (p Person) SetAge(age int) {
	p.Age = age
}

// 修改年龄
func (p *Person) SetAge2(age int) {
	p.Age = age
}
