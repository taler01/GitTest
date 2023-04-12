package main

import "fmt"

func main() {
	var array = [3]int{1, 2, 3}
	for i, v := range array {
		fmt.Println("%d %d\n", i, v)
	}
	a := []int{1, 2, 3, 4, 5}
	b := a[:3]
	fmt.Println(cap(a))
	fmt.Println(len(a))
	fmt.Println(cap(b))
	fmt.Println(len(b))

	b = append(b, 1, 2, 3)
	fmt.Println(b)
	//当容量不够时，通过翻倍来解决（容量和长度不是一个概念）

	fmt.Println(cap(a))
	fmt.Println(len(a))
	fmt.Println(cap(b))
	fmt.Println(len(b))
}

//func tete() {
//	var t int
//	t = 2
//	fmt.Println("数值是：#{t}")
//}
