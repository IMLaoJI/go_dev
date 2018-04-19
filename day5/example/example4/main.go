package main

import "fmt"

type integer int

type Student struct {
	Number int
}

type Stu Student //alias
//make 用来创建map、slice、channel
//new用来创建值类型 用户自定义的结构体类型也是值类型
func main() {

	var i integer = 1000
	var j int = 100

	j = int(i)
	fmt.Println(j)

	var a Student
	a = Student{30}

	var b Stu
	a = Student(b)
	fmt.Println(a)

}
