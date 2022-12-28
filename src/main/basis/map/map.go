package main

import (
	"fmt"
	"sort"
)

/*
*
格式 : map[KeyType]ValueType
*/
func main() {
	makeMap()
	initMapData()
	checkValueIsExits()
	loopKeyOrValue()
	deleteKey()
	sortKeyLoop()
}

/*
*
map 初始化默认是Nil,需要通过 Make 进行内存分配
*/
func makeMap() {
	fmt.Println("------------------------------------")
	// key : string
	// value: string
	// 容量cap : 10[非必填]
	testMap := make(map[string]string, 10)
	testMap["a"] = "a"
	testMap["b"] = "b"
	for key, value := range testMap {
		fmt.Printf("key:%s,value:%s", key, value)
		fmt.Println()
	}
	fmt.Println(len(testMap))
}

/*
*
创建Map并给初始值
*/
func initMapData() map[string]string {
	fmt.Println("------------------------------------")
	testMap := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}
	for k, v := range testMap {
		fmt.Printf("key:%s,value:%s", k, v)
		fmt.Println()
	}
	fmt.Println(len(testMap))
	return testMap
}

/*
*
判断值是否存在
*/
func checkValueIsExits() {
	fmt.Println("------------------------------------")
	testMap := initMapData()
	if testMap == nil {
		fmt.Println("map为空")
	} else {
		fmt.Println("非空Map")
		v, ok := testMap["a"]
		if ok {
			fmt.Printf("d存在 = %s", v)
		} else {
			fmt.Println("d不存在 ")
		}
	}
}

/*
*
遍历Key或Value
*/
func loopKeyOrValue() {
	fmt.Println("------------------------------------")
	testMap := initMapData()
	if testMap == nil {
		return
	}
	for key := range testMap {
		fmt.Printf("key = %s", key)
		fmt.Println()
	}
	for _, value := range testMap {
		fmt.Printf("value = %s", value)
		fmt.Println()
	}
}

/*
*
删除某个Key
*/
func deleteKey() {
	fmt.Println("---------------deleteKey---------------------")
	testKey := initMapData()
	if testKey == nil {
		return
	}
	fmt.Println("---------------删除前---------------------")
	fmt.Println(testKey)
	delete(testKey, "a")
	fmt.Println("---------------删除后---------------------")
	fmt.Println(testKey)
}

/*
*
先对Key进行排序再按Key顺序列出value
*/
func sortKeyLoop() {
	fmt.Println("------------------------------------")
	// 创建一个map并给值
	tMap := make(map[string]string)
	tMap["d"] = "d"
	tMap["s"] = "s"
	tMap["b"] = "d"
	tMap["b"] = "b"
	tMap["e"] = "e"
	tMap["a"] = "a"
	// 创建一个切片
	keySlice := make([]string, 0)
	for key := range tMap {
		keySlice = append(keySlice, key)
	}
	if keySlice == nil {
		return
	}
	fmt.Println("---------------Key切片排序前---------------------")
	for index, key := range keySlice {
		fmt.Printf("index=%d,key=%s", index, key)
		fmt.Println()
	}
	sort.Strings(keySlice)
	fmt.Println("---------------Key切片排序后---------------------")
	for index, key := range keySlice {
		fmt.Printf("index=%d,key=%s", index, key)
		fmt.Println()
	}
	fmt.Println("---------------map排序后---------------------")
	for _, key := range keySlice {
		fmt.Printf("key:%s,value:%s", key, tMap[key])
		fmt.Println()
	}
}
