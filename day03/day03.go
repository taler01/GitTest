package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)
	ages["test01"] = 1
	ages["test02"] = 2
	fmt.Println(ages["test01"])
	// 另一种创建hash
	OtherInitMap()
	// 删除元素
	delete(ages, "test01")
	fmt.Println(ages["test01"])
	type test struct {
		name     string
		age      int
		salary   int
		manageID int
	}
}
func OtherInitMap() {
	ages := map[string]int{
		"alice": 31,
		"bob":   32,
	}
	fmt.Println(ages["bob"])
	for name, age := range ages {
		fmt.Println("%s\t%d\n", name, age)
	}
	println("---------------------------------------------")
	for name1, age1 := range ages {
		println(name1, age1)
	}
}
