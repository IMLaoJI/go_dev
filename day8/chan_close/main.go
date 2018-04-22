package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	//close(ch)
	go func(){for {
		var b int
		b = <-ch
		fmt.Println("sss")
		//if ok == false {
		//	fmt.Println("chan is close")
		//	break
		//}
		fmt.Println(b)
	}}()
	ch <- 3333
	time.Sleep(1999*time.Second)
}
