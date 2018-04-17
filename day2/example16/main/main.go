package main

import "fmt"
import "strconv"
//https://go-zh.org/pkg/fmt/
//https://studygolang.com/articles/11119

//两种int类型转化成string类型的方法
func intTostring(n int) (string,string) {
	str1 := strconv.Itoa(n)
	str2 := fmt.Sprintf("%d", n)
	return str1,str2
}
func main() {
	var a int = 100
	var b bool
	c := 'a'

	fmt.Printf("%+v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("90%%\n")
	fmt.Printf("%t\n", b)
	fmt.Printf("%b\n", 100)
	fmt.Printf("%f\n", 199.22)
	fmt.Printf("%q\n", "this is a test")
	fmt.Printf("%x\n", 39839333)
	fmt.Printf("%p\n", &a)

	str := fmt.Sprintf("a=%d", a)
	fmt.Printf("%q\n", str)
	s1, s2 := intTostring(a)
	fmt.Printf("type(s1) = %T, type(s2) = %T",s1,s2)

}
