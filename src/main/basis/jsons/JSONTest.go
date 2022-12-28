package main

import (
	"encoding/json"
	"fmt"
	"go.leaving.com/src/main/basis/structs"
)

func main() {
	person := &structs.Person{
		"jianghai",
		22,
		"陕西省西安",
	}
	// 对象转JSON
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("json marshal failed")
	}
	fmt.Printf("json:%s\n", data)
	// JSON转对象
	newPerson := &structs.Person{}
	err = json.Unmarshal(data, newPerson)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", newPerson)
}
