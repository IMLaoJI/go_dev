package main

import(
	"time"
	"go_dev/day1/goroute/goroute"
)

func main() {

	for i := 0; i < 100; i++ {
		go goroute.Test_goroute(i)
	}

	time.Sleep(time.Second)
}