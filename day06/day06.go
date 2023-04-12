package main

import (
	"fmt"
	"math"
)

const globalValue int = 1

type testStruct struct {
	n int
}

func main() {
	age := map[string]int{
		"tom":    32,
		"longzi": 34,
	}
	age["alice"] = 72
	fmt.Println("value:", age["tom"])
	for k, v := range age {
		fmt.Println(k, v)
	}
	delete(age, "tom")
	for k, v := range age {
		fmt.Println(k, v)
	}
	for k, v := range age {
		fmt.Printf("\nk:%v v:%v\n", k, v)
	}
	var names []string
	for k, _ := range age {
		names = append(names, k)
	}
	for i, v := range names {
		fmt.Printf("\nindex:%v  value:%v\n", i, v)
	}
	x := MethodTest(3, 4)
	fmt.Println(x)
	a := testStruct{0}
	fmt.Println("main函数中a的值：", a.n)
	test(a)
	Test(&a)
	fmt.Println("经过取地址处理后的结构体的数据：", a.n)
	// 引用类型只有 slice(切片)、channel(管道)和map;
	// 剩下的都是值类型；(值类型需要对其进行取地址操作才能修改原值)
	/*
		引用类型的变量，初始化方式也不一样：
		slice类型，用make，new，{}都可以
		map类型，用make，new，{}都可以
		channel类型，只能用make活着new初始化
		//map可以用{}，make，new三种方式初始化
	*/
	/*
		make和new的区别
		make返回的是对象。
		对值类型对象的更改，不会影响原始对象的值
		对引用类型对象的更改，会影响原始对象的值
		new返回的是对象的指针，对指针所在对象的更改，会影响指针指向的原始对象的值
	*/
	EE := AdrTest(3)
	fmt.Println(EE)
}
func MethodTest(x, y float64) (result float64) {
	result = math.Sqrt(x*x + y*y)
	fmt.Println(globalValue)
	return result
}
func test(x testStruct) {
	x.n = 111
	fmt.Println("未带*函数输出值：", x.n)
}
func Test(x *testStruct) {
	x.n = 111
	fmt.Println("带*函数输出值：", x.n)
}
func AdrTest(x int) (y int) {
	z := &x
	y = *z
	fmt.Println("经取地址赋值后，y:", y)
	return y
}
