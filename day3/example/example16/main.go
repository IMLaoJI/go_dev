package main

import "fmt"

var (
	result = func(a1 int, b1 int) int {
		return a1 + b1
	}
)

func test(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}

	return result(a, b)
}

//命名返回值的名字：
func add(a, b int) (c int) {
	c = a + b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a +b)/2
	return
}
/*	1. 当函数返回时，执行defer语句。因此，可以用来做资源清理
	2. 多个defer语句，按先进后出的方式执行
	3. defer语句中的变量，在defer声明时就决定了。
*/
func main() {
	fmt.Println(result(100, 200))

	var i int = 0
	defer fmt.Println(i)
	defer fmt.Println("second")

	i = 10
	fmt.Println(i)
	//_标识符，用来忽略返回值
	sum, _ := calc(100, 22)
	fmt.Println(sum)
}
