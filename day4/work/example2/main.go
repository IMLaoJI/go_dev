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
	//newa := append(a,1,2,3,4,5,4)
	//fmt.Println(cap(newa))
	//fmt.Println(len(newa))
	//fmt.Printf("%p\n",&newa)
	//test(3)
	slice := []int{10, 20, 30, 40}
	fmt.Printf("%p\n",&slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
	slice111 := []int{10, 20, 30, 40}
	newSlice := append(slice, slice111...)
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))
	//fmt.Printf("%p\n",&newSlice)
	//newSlice1 := append(newSlice, 50)
	//fmt.Printf("%p\n",&newSlice1)
	//fmt.Println(cap(newSlice1))
	//fmt.Println(len(newSlice1))
}
