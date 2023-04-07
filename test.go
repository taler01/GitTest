package main

import "fmt"

func main() {
	fmt.Println("hello")
	var s, t int
	// 注释符
	s = 1
	t = 2
	fmt.Println(s, t)
	// var 变量 类型 = 表达式 （其中“类型”或者“=表达式”两部分可以省略其中一个）；
	var test = "hello world"
	fmt.Println(test)
	//trest := "strong"
	//fmt.Println(trest)
	TestPrint("hello test")
	x := 1
	p := &x
	fmt.Println(p)  // 地址值
	fmt.Println(*p) // 地址对应的数值
	TestArray()
	var testArray = [...]int{1, 2, 3, 4} // 三个点是元素给定，不需要写出长度，并不是元素不缺定；
	var testArray1 = []int{1, 2, 3, 4}   // 元素不确定时，在中括号里面不需要写三个点
	fmt.Println(testArray)
	testArray1 = append(testArray1, 5, 6, 7)
	fmt.Println(testArray1)
	mod := modify()
	testmodify(mod)
}
func TestPrint(s string) {
	// 在函数内部，可以使用 “名字 ：= 表达式”形式定义变量；
	test := s
	fmt.Println(test)
}

// TestArray 数组测试
// 数组不可动态变化问题，一旦声明了，其长度就是固定的
func TestArray() {
	var arr_1 [5]int
	fmt.Println(arr_1)

	var arr_2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_2)

	arr_3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_3)

	arr_4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr_4)

	arr_5 := [5]int{0: 3, 1: 5, 4: 6}
	fmt.Println(arr_5)

	//二维数组
	var arr_6 = [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_6)

	arr_7 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_7)

	arr_8 := [...][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0: 3, 1: 5, 4: 6}}
	fmt.Println(arr_8)
}

func modify() [5]int {
	var arr_2 = [5]int{1, 2, 3, 4, 5}
	return arr_2
}

func testmodify(a [5]int) {
	a[1] = 20
	fmt.Println(a)
}
