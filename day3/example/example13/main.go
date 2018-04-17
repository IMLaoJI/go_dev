package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func test() {
	return
}

func main() {

	c := add
	sum := c(200, 300)
	fmt.Println(c)
	//if  c == add  {
	//	fmt.Println("c equal add")
	//}
		fmt.Println(sum)

	//str := "hello, world,中国"

	//for index, val := range str {
	//	fmt.Printf("index[%d] val[%c] len[%d]\n", index, val, len([]byte(string(val))))
	//}
	//
	//for index, val := range str {
	//	if index>2{
	//		continue
	//	}
	//	if (index > 3) {
	//		break
	//	}
	//	fmt.Printf("index[%d] val[%c] len[%d]\n", index, val, len([]byte(string(val))))
	//}

	//LABEL1:
	//	for i := 0; i <= 5; i++ {
	//		for j := 0; j <= 5; j++ {
	//			if j == 4 {
	//				continue LABEL1
	//			}
	//			fmt.Printf("i is: %d, and j is: %d\n", i, j)
	//		}
	//	}

	i := 0
	HERE:
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto HERE



}
