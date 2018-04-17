package main

import "fmt"

func Print(n int) {

	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
	var i =3
	for i > 0 {
		fmt.Println("i > 0")
	}

	for true{
		fmt.Println("i > 0")
	}

	for {
		fmt.Println("i > 0")
	}
}

func main() {
	Print(6)
}
