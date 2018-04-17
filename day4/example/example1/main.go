package main

import "fmt"
import (
	"errors"
)

func initConfig() (err error) {
	return errors.New("init config failed")
}

func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := initConfig()
	if err != nil {
		panic(err)
	}
	return
}
func new1(){
	var i int
	fmt.Println(i)

	j := new(int)
	*j = 100
	fmt.Println(*j)
}

func main() {
	//new1()
	//for {
	//	test()
	//	time.Sleep(time.Second)
	//}
	var b1 []int
	var c = []int{1,2,3}
	b1 = append(c,1,2,3)
	fmt.Println(b1)
	var a []int
	a = append(a, 10, 20, 383)
	a = append(a, a...)
	fmt.Println(a)

}
