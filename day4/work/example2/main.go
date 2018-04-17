package main

import "fmt"

func perfect(n int) bool {

	var sum int = 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return n == sum
}

func process(n int) {
	for i := 1; i < n+1; i++ {
		if perfect(i) {
			fmt.Println(i)
		}
	}
}
func test(n int)  {
	var a = make([]int,n)
	fmt.Println(a)

}

func main() {
	//var n int
	//fmt.Scanf("%d", &n)
	//process(n)
	//var a = make([]int, 3,4)
	//fmt.Printf("%p\n",&a)
	//fmt.Println(a)
	//fmt.Println(cap(a))
	//fmt.Println(len(a))
	//a= append(a,1,2,3,4,5,5,5,5,5,6,7,8)
	//fmt.Println(cap(a))
	//fmt.Println(len(a))
	//fmt.Printf("%p\n",&a)
	test(3)
}
