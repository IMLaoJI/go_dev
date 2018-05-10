package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

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
	"main",
}
var operator []string = []string{"+", "-", "*", "/", "<=", ">=", "<", ">", "=",
	"%", "+=", "-=", "/=", "*=", "!=", "&", "|", "<<", ">>",}

var division []string = []string{",", ";", ".", "(", ")", "{", "}", "[", "]", ":", "'",}
var sign []string = []string{"-", "=", "+", "<", ">", "*", "/", "!",}

func main() {
	file, err := os.Open(`H:\Go\Development\src\go_dev\wordAnalysis\te1st.txt`)
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		word := []rune{}
		word_list := []string{}
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v", err)
			break
		}
		str = strings.TrimSpace(str)
		runeArr := []rune(str)
		var flagstutus int = 0
		var flagstutus1 int = 0
		for index, v := range runeArr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				word = append(word, v)
			case v == '_':
				word = append(word, v)
			case v == '\r' || v == '\n':

			case v >= '0' && v <= '9':
				word = append(word, v)
			default:
				flagstutus1 = 0
				if len(word) != 0 {
					word_list = append(word_list, string(word))
					word = []rune{}
				}

				if v != ' ' {
					wordtmp := []rune{}
					for _, val := range sign {
						wordtmp = []rune{}
						if string(v) == val {
							flagstutus1 = 1
							for _, val := range sign {
								if string(runeArr[index-1]) == val {
									wordtmp = append(wordtmp, runeArr[index-1])
									wordtmp = append(wordtmp, v)
									flagstutus = 1
									word_list = word_list[0:len(word_list)-1]
									break
								}
							}
							if flagstutus == 1 {
								break
							}
						}
						if flagstutus1 == 1 {
							if flagstutus == 0 {
								wordtmp = append(wordtmp, v)
							}

						}

					}
					if flagstutus1 == 0 {
						wordtmp = append(wordtmp, v)
					}

					word_list = append(word_list, string(wordtmp))
				}

			}
		}
		//fmt.Println(word_list)
		dealtypes(word_list)

	}

}
func dealtypes(word_list []string) {

	for _, v := range word_list {
		flag1:=1
		flag2:=1
		flag3:=1
		if v >= "0" && v <= "9" {
			fmt.Printf("(3,%s)\n", v)
			continue
		}
		for _, v1 := range KEY_SET {
			if v1 == v {
				fmt.Printf("(1,%s)\n", v)
				flag1=0
				break
			}
		}
		if flag1==0 {
			continue
		}
		for  _, v2 := range operator {
			if v2 == v {
				fmt.Printf("(4,%s)\n", v)
				flag2=0

				break
			}
		}
		if flag2==0 {
			continue
		}
		for  _, v3 := range division {
			if v3 == v {
				fmt.Printf("(5,%s)\n", v)
				flag3=0
				break
			}
		}
		if flag3==0 {
			continue
		}

		fmt.Printf("(2,%s)\n", v)
	}
}
