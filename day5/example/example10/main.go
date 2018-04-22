package main

import "fmt"

type People struct {
	name string
	age  int
}

type Test interface {
	Print(string)string
	Sleep()
}

type Student struct {
	name  string
	age   int
	score int
}

func (p *Student) Print(string string)string {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("score:", p.score)
	return ""
}

func (p *Student) Sleep() {
	fmt.Println("student sleep")
}

func (people People) Print() {
	fmt.Println("name:", people.name)
	fmt.Println("age:", people.age)
}

func (p People) Sleep() {
	fmt.Println("people sleep")
}
//Interface类型可以定义一组方法，但是这些不需要实现。并且interface不能
//包含任何变量。
func main() {

	var t Test
	fmt.Println(t)
	//t.Print()

	var stu Student = Student{
		name:  "stu1",
		age:   20,
		score: 200,
	}

	t = &stu
	t.Print("ss")
	t.Sleep()

	//var people People = People{
	//	name: "people",
	//	age:  100,
	//}
	//
	//t = people
	//t.Print()
	//t.Sleep()
	//
	//fmt.Println("t:", t)
}
