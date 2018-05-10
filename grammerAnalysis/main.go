package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	fmt.Printf("递归下降分析程序，编制人：%s，%s，%s\n","姬龙龙","201521122082","软件1511")
	fmt.Print("输入一以#结束的符号串(包括+-*/()i#)：")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	trimmedInput := strings.Trim(input, " \n")
	trimmedInput1 := []rune(trimmedInput)
	str :=deal(trimmedInput1)
	fmt.Printf("输出结果：%s为%s",trimmedInput,str)
}
func deal(input []rune) string {
	tmp :=input
	input = append(input[:1],input[1:len(input)-1]...)
	judge := dealE(input)
	if judge&& tmp[len(tmp)-1]=='#'{
		return "合法符号串"
	} else {
		return "非法的符号串"
	}
}

func dealE(input []rune) bool {
	if dealT(input) && dealG(input) {

		return true
	}
	return false
}

func dealG(input []rune) bool {
	if input[0] == '+' {
		input = append(input[:0],input[1:]...)
		if dealT(input) && dealG(input) {
			return true
		}
		return false
	} else if input[0] == '-' {
		input = append(input[:0],input[1:]...)
		if dealT(input) && dealG(input) {
			return true
		}
		return false
	}else {
		return true
	}
}
func dealT(input []rune) bool {
	if dealF(input) && dealS(input) {
		return true
	}
	return false
}

func dealF(input []rune) bool {
	if input[0] == '(' {
		input = append(input[:0],input[1:]...)
		if dealE(input) {
			if input[0] == ')' {
				input = append(input[:0],input[1:]...)
				return true
			}
		}
	} else if input[0] == 'i' {
		input = append(input[:0],input[1:]...)
		return true
	}
	return false
}
func dealS(input []rune) bool {
	if input[0] == '*' {
		input = append(input[:0],input[1:]...)
		if dealF(input) && dealS(input) {
			return true
		}
		return false
	} else if input[0] == '/' {
		input = append(input[:0],input[1:]...)
		if dealF(input) && dealS(input) {
			return true
		}
		return false
	}else {
		return true
	}
}
