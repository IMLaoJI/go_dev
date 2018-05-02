package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
)

func readfile1(){
	f, err := os.Open("H:\\Go\\Development\\src\\go_dev\\wordAnalysis\\te1st.txt")
	if err != nil {
		panic("open failed!")
	}
	defer f.Close()
	buff := make([]byte, 1024)
	for n, err := f.Read(buff); err == nil; n, err = f.Read(buff) {
		fmt.Print(string(buff[:n]))
	}
	if err != nil {
		panic(fmt.Sprintf("Read occurs error: %s", err))
	}
}

func readfile2(){
	buff, err := ioutil.ReadFile("H:\\Go\\Development\\src\\go_dev\\wordAnalysis\\te1st.txt")
	if err !=nil{
		panic("open file failed")
	}
	fmt.Println(string(buff))

}

func readfile3(filename string){
	file, err := os.Open(filename)
	if err != nil {
		panic("open failed!")
	}
	defer file.Close()
	b := bufio.NewReader(file)
	line, err := b.ReadString('\n')
	for ; err == nil; line, err = b.ReadString('\n'){
		fmt.Print(line)
	}
	if err == io.EOF {
		//fmt.Println("test")
		fmt.Print(line)
	}else {
		panic("read occur error")
	}

}

func judgeKeywords(char string) int{
	keywords :=[]string{"if","int","for","while","do","return","break","continue"}
	for i := 0; i<=7;i++  {
		if strings.Compare(char,keywords[i]) ==0{
			return i+1
		}
	}
	return 0
}

func getChar(line string) rune {
	return line
}
func scanner(line string) {
	state := 0
	ch := " "
	pos := 0
	for ;ch != "\0";{
		switch state {
		case 0:
			ch = getChar(line)
		}
	}
}