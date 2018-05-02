package main

//import "fmt"

const KEY_NUM int = 11

var KEY_SET []string = []string{
	"for",
	"while",
	"do",
	"continue",
	"if",
	"else",
	"char",
	"int",
	"double",
	"float",
	"return",
}

const filename string = "H:\\Go\\Development\\src\\go_dev\\wordAnalysis\\te1st.txt"
func main() {
	readfile3(filename)
	//fmt.Println(KEY_SET)
}
