package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var sum int
var NODE, ENODE string

var n []Edge

func initWF() {
	fmt.Print("请输入上下文无关文法的总规则数：")
	inputReader := bufio.NewReader(os.Stdin)
	num, _ := inputReader.ReadString('\n')
	num = strings.Trim(num, " \n")
	numRule, err := strconv.Atoi(num)
	sum = numRule
	if err != nil {
		fmt.Println("输入的不是数字哦")
		return
	}
	n = make([]Edge,sum)
	fmt.Println("请输入具体规则（格式：左部 右部，@为空）：")
	for j := 0; j < numRule; j++ {
		WF, _ := inputReader.ReadString('\n')
		WF = strings.Trim(WF, " \n")
		split := strings.Split(WF, " ")
		//n[j] = Edge{
		//	left:     "",
		//	right:    "",
		//	rlen:     0,
		//	first:    "",
		//	follow:   "",
		//	myselect: "",
		//}
		n[j].left = split[0]
		n[j].right = split[1]
		n[j].rlen = len(n[j].right)
		if !strings.Contains(NODE, n[j].left) {
			NODE += n[j].left
		}
	}
}

func main() {
	fmt.Printf("递归下降分析程序，编制人：%s，%s，%s\n", "姬龙龙", "201521122082", "软件1511")
	//初始化文法
	initWF()
	//处理符号集合
	dealFH()
	//处理first集
	dealFirst()
	//处理follow集
	dealfollow()
	//处理select集
	flag := dealselect()
	//输出文法的first集和follow集 并且判断是否为LL文法
	drawone(flag)
	//输出预测分析表
	yc := drawtwo()
	//处理分析串
	dealAnalysis(yc)

}
func dealAnalysis(yc [][]string) {
	var chuan, fenxi, fchuan string
	inputReader := bufio.NewReader(os.Stdin)
	for ; ; {
		fmt.Print("\n请输入符号串：")
		chuan, _ = inputReader.ReadString('\n')
		chuan = strings.Trim(chuan, " \n")
		temp := []rune(chuan)
		if len(temp) == 0 {
			fmt.Print("输入错误!!!")
			continue
		} else {
			break
		}
	}
	fchuan = chuan
	fenxi = "#"
	fenxi += string(NODE[0])
	var i int = 0
	fmt.Println("预测分析过程如下：")
	fmt.Print("步骤")
	drawFH(7, " ")
	fmt.Print("分析栈")
	drawFH(10, " ")
	fmt.Print("剩余输入串")
	drawFH(8, " ")
	fmt.Print("推导所用产生式或匹配")
	if match(&chuan, &fenxi, yc, &i) == 1 {
		fmt.Print("\n输入串", fchuan, "是该文法的句子\n")
	} else {
		fmt.Print("\n输入串", fchuan, "不是该文法的句子\n")
	}
}
func match(chuan, fenxi *string, yc [][]string, b *int) int {
	var ch, a rune
	var x, i, j, k int
	chu := []rune(*chuan)
	fen := []rune(*fenxi)
	*b++
	fmt.Print("\n", " ", *b)
	if *b > 9 {
		drawFH(8, " ")
	} else {
		drawFH(9, " ")
	}
	fmt.Print(*fenxi)
	drawFH(26-len(*chuan)-len(*fenxi), " ");
	fmt.Print(*chuan)
	drawFH(10, " ")
	a = chu[0]
	ch = fen[len(fen)-1]
	x = strings.Index(ENODE, string(ch))
	if x < len(ENODE) && x > -1 {
		if ch == a {
			fen = erase(fen, len(fen)-1)
			*fenxi = string(fen)
			chu = erase(chu, 0)
			*chuan = string(chu)
			fmt.Print("'", string(a), "'匹配")
			if match(chuan, fenxi, yc, b) == 1 {
				return 1
			} else {
				return 0
			}
		}
		return 0
	} else {
		if string(ch) == "#" {
			if ch == a {
				fmt.Println("分析成功")
				return 1
			} else {
				return 0
			}
		} else {
			if string(ch) == "@" {
				fen = erase(fen, len(fen)-1)
				*fenxi = string(fen)
				if match(chuan, fenxi, yc, b) == 1 {
					return 1
				} else {
					return 0
				}
			} else {
				i = strings.Index(NODE, string(ch))
				if string(a) == "#" {
					x = strings.Index(ENODE, "@")
					if x < len(ENODE) && x > -1 {
						j = len(ENODE) - 1
					} else {
						j = len(ENODE)
					}
				} else {
					j = strings.Index(ENODE, string(a))
				}
				yctem := []rune(yc[i][j])
				if len(yctem) > 0 {
					fmt.Print(string(NODE[i]), "->", yc[i][j])
					fen = erase(fen, len(fen)-1)
					*fenxi = string(fen)
					for k = len(yc[i][j]) - 1; k > -1; k-- {
						if string(yc[i][j][k]) != "@" {
							*fenxi += string(yc[i][j][k])
						}
					}
					if match(chuan, fenxi, yc, b) == 1 {
						return 1
					} else {
						return 0
					}
				} else {
					return 0
				}
			}
		}
	}

}
func erase(tem []rune, index int) []rune {
	tem = append(tem[:index], tem[index+1:len(tem)]...)
	return tem
}
func drawtwo() [][]string {
	var LOOP_COUNT = len(NODE)
	var LOOP_NUM = 50
	fmt.Println("\n\n预测分析表如下")
	//var tem  =len(NODE)
	var yc [][]string
	for i := 0; i < LOOP_COUNT; i++ {
		sl := make([]string, 0, LOOP_NUM)
		for j := 0; j < LOOP_NUM; j++ {
			sl = append(sl, "")
		}
		yc = append(yc, sl)
	}
	var i, j, k int
	var flag int
	for i = 0; i < len(ENODE); i++ {
		if string(ENODE[i]) != "@" {
			drawFH(10, " ")
			fmt.Print(string(ENODE[i]))
		}
	}
	drawFH(10, " ")
	fmt.Println("#")
	var x int
	for i = 0; i < len(NODE); i++ {
		drawFH(4, " ")
		fmt.Print(string(NODE[i]))
		drawFH(5, " ")
		for k = 0; k < len(ENODE); k++ {
			flag = 1
			for j = 0; j < sum; j++ {
				if NODE[i] == n[j].getlf()[0] {
					x = strings.Index(n[j].getselect(), string(ENODE[k]))
					if x < len(n[j].getselect()) && x > -1 {
						fmt.Print("->", n[j].getrg())
						yc[i][k] = n[j].getrg()
						drawFH(9-n[j].getrlen(), " ")
						flag = 0
					}
					x = strings.Index(n[j].getselect(), "#")
					if (k == len(ENODE)-1 && x < len(n[j].getselect()) && x > -1) {
						fmt.Print("->", n[j].getrg())
						yc[i][j] = n[j].getrg()
					}
				}
			}
			if flag == 1 && string(ENODE[k]) != "@" {
				drawFH(11, " ")
			}
		}
		fmt.Println()
	}
	return yc

}
func drawone(flag int) {
	var i, j int
	fmt.Printf("\n非终结符")
	drawFH(sum+2, " ")
	fmt.Print("First")
	drawFH(sum+2, " ")
	fmt.Print("Follow\n")
	drawFH(5+sum, "_*_")
	fmt.Println()
	for i = 0; i < len(NODE); i++ {
		for j = 0; j < sum; j++ {
			if NODE[i] == n[j].getlf()[0] {
				drawFH(3, " ")
				fmt.Print(string(NODE[i]))
				drawFH(sum+4, " ")
				drawSet(n[j].getfirst())
				drawFH(sum+4-2*len(n[j].getfirst()), " ")
				drawSet(n[j].getfollow())
				fmt.Println()
				break
			}
		}
	}

	drawFH(5+sum, "_*_")
	fmt.Print("\n\n判定结论：		")
	if flag == 1 {
		fmt.Println("该文法不是LL(1)文法!")
	} else {
		fmt.Println("该文法是LL(1)文法!")
	}
}
func drawSet(p string) {
	var i int
	if len(p) == 0 {
		return
	}
	fmt.Print("{")
	for i = 0; i < len(p)-1; i++ {
		fmt.Print(string(p[i]), ",")
	}
	fmt.Print(string(p[i]), "}")
}
func drawFH(total int, str string) {
	var i int
	for i = 0; i < total; i++ {
		fmt.Print(str)
	}
}
func dealfollow() {
	var k, i, j int
	for k = 0; k < sum; k++ {
		for i = 0; i < sum; i++ {
			if n[i].getlf() == n[0].getlf() {
				n[i].newfollow("#")
			}
			follow(&n[i], n, i)
		}
		for i = 0; i < sum; i++ {
			for j = 0; j < sum; j++ {
				if strings.Index(n[j].getrg(), n[i].getlf()) == n[j].getrlen()-1 {
					n[i].newfollow(n[j].getfollow())
				}
			}
		}
	}
}
func follow(e *Edge, n []Edge, x int) {
	var i, j, k, s int
	var str string
	for i = 0; i < e.getrlen(); i++ {
		s = strings.Index(NODE, string(e.getrg()[i]))
		if s < len(NODE) && s > -1 { //是非终结符
			if i < e.getrlen()-1 { //不在最右
				for j = 0; j < sum; j++ {
					if strings.Index(n[j].getlf(), string(e.getrg()[i])) == 0 {
						if strings.Contains(NODE, string(e.getrg()[i+1])) {
							for k = 0; k < sum; k++ {
								if strings.Index(n[k].getlf(), string(e.getrg()[i+1])) == 0 {
									n[j].newfollow(n[k].getfirst())
									if strings.Index(n[k].getfirst(), "@") < len(n[k].getfirst()) {
										n[j].newfollow(e.getfollow())
									}
								}
							}
						} else {
							str = ""
							str += string(e.getrg()[i+1])
							n[j].newfollow(str)
						}
					}
				}
			}
		}
	}
}
func dealselect() int {
	var i, j, k int
	var str string
	var flag int
	for i = 0; i < sum; i++ {
		myselect(&n[i], n)
	}
	for i = 0; i < len(NODE); i++ {
		str = ""
		for j = 0; j < sum; j++ {
			if n[j].getlf()[0] == NODE[i] {
				if len(str) == 0 {
					str = n[j].getselect()
				} else {
					for k = 0; k < len(n[j].getselect()); k++ {
						if strings.Contains(str, string(n[j].getselect()[k])) {
							flag = 1
							break
						}
					}
				}
			}
		}
	}
	return flag
}
func myselect(e *Edge, n []Edge) {
	var i, j int
	if strings.Contains(ENODE, e.getro()) {
		e.newselect(e.getro())
		if e.getro() == "@" {
			e.newselect(e.getfollow())
		}
	} else {
		for i = 0; i < e.getrlen(); i++ {
			for j = 0; j < sum; j++ {
				if e.getrg()[i] == n[j].getlf()[0] {
					e.newselect(n[j].getfirst())
					if !strings.Contains(n[j].getfirst(), "@") {
						return
					}
				}
			}
		}
	}
}
func first(e *Edge, n []Edge, x int) {
	var i, j int
	for j = 0; j < sum; j++ {
		if e.getlf() == n[j].getlf() {
			if strings.Contains(NODE, n[j].getro()) {
				for i = 0; i < sum; i++ {
					if n[i].getlf() == n[j].getro() {
						first(&n[i], n, x)
					}
				}
			} else {
				n[x].newfirst(n[j].getro())
			}
		}
	}
}
func dealFirst() {
	var i, j, k int
	for i = 0; i < sum; i++ {
		first(&n[i], n, i)
	}
	for i = 0; i < sum; i++ {
		if strings.Contains(n[i].getfirst(), "@") {
			if strings.Contains(NODE, n[i].getro()) {
				for k = 1; k < n[i].getrlen(); k++ {
					if strings.Contains(NODE, string(n[i].getrg()[k])) {
						for j = 0; j < sum; j++ {
							if n[i].getrg()[k] == n[j].getlf()[0] {
								n[i].newfirst(n[j].getfirst())
								break
							}
						}
						if !strings.Contains(n[j].getfirst(), "@") {
							n[i].delfirst()
							break
						}
					}
				}
			}
		}
	}

}
func dealFH() {
	for i := 0; i < sum; i++ {
		for j := 0; j < n[i].getrlen(); j++ {
			str := n[i].getrg()
			strtemp := string(str[j])
			if !strings.Contains(NODE, strtemp) && !strings.Contains(ENODE, strtemp) {
				ENODE += strtemp
			}
		}
	}
}
