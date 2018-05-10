package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"go_dev/datastructure/stack"
	"strconv"
)

var output []rune
var sign []string = []string{"-", "=", "+", "<", ">", "*", "/", "!",}
func main() {
	fmt.Printf("递归下降分析程序，编制人：%s，%s，%s\n", "姬龙龙", "201521122082", "软件1511")
	fmt.Print("输入一以#结束的中缀表达式(包括+—*/（）数字#)：")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	trimmedInput := strings.Trim(input, " \n")
	//dealstr := dealstr(trimmedInput)
	//fmt.Println(dealstr)

	trimmedInput1 := []rune(trimmedInput)
	str := infix_to_suffix(trimmedInput1)
	fmt.Printf("前缀转后缀：%s = %s\n", trimmedInput, str)
	anwer, err := deal(output)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("后缀的值：%s = %d\n", str, anwer)
}
func deal(out []rune) (int, error) {
	is_num := false
	stack := &stack.ItemStackInt{}
	itemStack := stack.New()
	for _, item := range out {
		if item >= '0' && item <= '9' {
			is_num = true
		} else {
			is_num = false
		}
		if is_num {
			temp, _ := strconv.Atoi(string(item))
			itemStack.Push(temp)
		} else {
			flag := 0
			switch item {
			case '+':
				flag = 1
			case '-':
				flag = 2
			case '*':
				flag = 3
			case '/':
				flag = 4
			default:
				flag = 5
			}
			if flag == 1 {
				a := *itemStack.Pop()
				b := *itemStack.Pop()
				itemStack.Push(a + b)
			} else if flag == 2 {
				a := *itemStack.Pop()
				b := *itemStack.Pop()
				itemStack.Push(b - a)
			} else if flag == 3 {
				a := *itemStack.Pop()
				b := *itemStack.Pop()
				itemStack.Push(a*b)
			} else if flag == 4 {
				a := *itemStack.Pop()
				b := *itemStack.Pop()
				if a == 0 {
					err := fmt.Errorf("除数为0")
					return 0, err
				}
				itemStack.Push(b / a)
			} else {
			}

		}

	}
	//if itemStack.Size() == 1 {
	//	return int(*itemStack.Peek()), nil
	//} else {
	//	size := itemStack.Size()
	//	res := ""
	//	for j := 0; j < size; j++ {
	//		res = fmt.Sprintf("%s%s", res, int(*itemStack.Pop()))
	//	}
	//	res1, _ := strconv.Atoi(res)
	//	return res1, nil
	//}
	return *itemStack.Peek(),nil

}
func infix_to_suffix(exp []rune) string {
	exp = append(exp[:1],exp[1:len(exp)-1]...)
	output = []rune{}
	stack := &stack.ItemStack{}
	itemStack := stack.New()
	if len(exp) == 0 {
		return "输入错误"
	}
	lengh := len(exp)
	for j := 0; j < lengh; j++ {
		flag := 0
		item := exp[j]
		switch item {
		case '+':
			fallthrough
		case '-':
			flag = 1
		case '*':
			fallthrough
		case '/':
			flag = 2
		case '(':
			flag = 3
		case ')':
			flag = 4
		default:
			flag = 5
		}
		if flag == 1 || flag == 2 {
			get_oper(item, flag, itemStack)
		} else if flag == 3 {
			itemStack.Push(item)
		} else if flag == 4 {
			get_parent(itemStack)
		} else if flag == 5 {
			output = append(output, item)
		}

	}
	for ; ; {
		if !itemStack.IsEmpty() {
			output = append(output, *itemStack.Pop())
			continue
		}
		break
	}
	res := []string{}
	for _, val := range output {
		res = append(res, string(val))
	}
	return strings.Join(res, " ")

}

func get_oper(opthis rune, prec1 int, stack *stack.ItemStack) {
	for ; stack.IsEmpty() != true; {
		optop := *stack.Pop()
		if optop == '(' {
			stack.Push(optop)
			break
		} else
		{
			prec2 := 0
			if optop == '+' || optop == '-' {
				prec2 = 1
			} else {
				prec2 = 2
			}
			if prec2 < prec1 {
				stack.Push(optop)
				break
			} else {
				output = append(output, optop)
			}
		}
	}
	stack.Push(opthis)
}
func get_parent(stack *stack.ItemStack) {
	for ; stack.IsEmpty() != true; {
		item := *stack.Pop()
		if item == '(' {
			break
		} else
		{
			output = append(output, item)
		}
	}
}
func dealstr(str string) []string  {
	word := []rune{}
	word_list := []string{}
	runeArr := []rune(str)
	for _, v := range runeArr {
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
			if len(word) != 0 {
				word_list = append(word_list, string(word))
				word = []rune{}
			}
				word_list = append(word_list, string(v))

		}

	}
	if len(word) != 0 {
		word_list = append(word_list, string(word))
		word = []rune{}
	}
	//fmt.Println(word_list)
	return word_list
}
