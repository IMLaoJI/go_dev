package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name string
	Age  int
}

func (c *Car) Set(name string, age int) {
	c.Name = name
	c.Age = age
}

type Car2 struct {
	Name string
}

type Train struct {
	Car
	Car2
	createTime time.Time
	int
}

func (t *Train) Set(age int) {
	t.int = age
}

func main() {
	var train Train
	train.int = 300
	train.Car.Set("huas", 100)

	train.Car.Name = "test"
	fmt.Println(train)
	a :="SSS"
	//array2 := [3]*string{} //等价于array2 :=new([3]*string)
	array2 := [3]*string{new(string), new(string), new(string)}
	fmt.Println(array2)
	array3 := new([3]*string)
	array3[0] = &a
	fmt.Println(*array3[0])
}
