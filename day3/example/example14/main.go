package main

import "fmt"

type op_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func operator(op op_func, a, b int) int {

	return op(a, b)
}
//注意2：map、slice、chan、指针、interface默认以引用的方式传递
func main() {
	var a, b int
	fmt.Println(add(a, b))

	var c op_func
	c = add
	fmt.Println(add)
	fmt.Println(c)

	sum := operator(c, 100, 200)
	fmt.Println(sum)
}
