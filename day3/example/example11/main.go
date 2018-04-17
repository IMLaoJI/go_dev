package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	//var n int
	//n = rand.Intn(100)
	//
	//if n == 1 {
	//	fmt.Printf("keyide")
	//}else {
	//	fmt.Println("shawanyi")
	//}
	//
	//if n == 1 {
	//	fmt.Printf("keyide")
	//}else if n == 3 {
	//	fmt.Printf("keyide")
	//}else {
	//	fmt.Println("d")
	//}
	//
	//for {
	//	var input int
	//	fmt.Scanf("%d\n", &input)
	//	flag := false
	//	switch {
	//	case input == n:
	//		fmt.Println("you are right")
	//		flag = true
	//	case input > n:
	//		fmt.Println("bigger")
	//	case input < n:
	//		fmt.Println("less")
	//	}
	//
	//	if flag {
	//		break
	//	}
	//}
	test1()
	test2()
	test4()
	test5()
}

func test1()  {
	var i = 0
	switch i{
	case 0:
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")

	}

}

func test2()  {
	var i = 0
	switch i{
	case 0:
		fallthrough
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")

	}

}

//从本例可以看出：switch从第一个expr为true的case开始执行，
// 如果case带有fallthrough，程序会继续执行下一条case,不会再判断下一条case的expr是否为true,
// 如果之后的case都有fallthrough,default则会被执行
//func test3() {
//	switch {
//	case false:
//		fmt.Println("The integer was <= 4")
//		fallthrough
//	case true:
//		fmt.Println("The integer was <= 5")
//		fallthrough
//	case false:
//		fmt.Println("The integer was <= 6")
//		fallthrough
//	case true:
//		fmt.Println("The integer was <= 7")
//	case false:
//		fmt.Println("The integer was <= 8")
//		fallthrough
//	default:
//		fmt.Println("default case")
//	}
//}

func test4()  {
	var i = 0
	switch i {
	case 0, 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
}

func test5()  {
	var i = 0
	switch {
	case i > 0 && i < 10:
		fmt.Println("i>0 and i<10")
	case i > 10 && i<20:
		fmt.Println("i>10 and i<20")
	default:
		fmt.Println("def")
	}
	
}

//func test6() {
//	switch i := 0 {
//	case i > 0 && i < 10:
//		fmt.Println("i>0 and i<10")
//	case i > 10 && i<20:
//		fmt.Println("i>10 and i<20")
//	default:
//		fmt.Println("def")
//	}
//}