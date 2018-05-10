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

var n [8]Edge = [8]Edge{}

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
	fmt.Println("请输入具体规则（格式：左部 右部，@为空）：")
	for j := 0; j < numRule; j++ {
		WF, _ := inputReader.ReadString('\n')
		WF = strings.Trim(WF, " \n")
		split := strings.Split(WF, " ")
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
	drawtwo()
	fmt.Println(ENODE)
	fmt.Println(NODE)

}
func drawtwo() {
	fmt.Println("\n\n预测分析表如下")
	yc := [len(NODE)][50]string{}
	var i, j, k int
	var flag bool
	for i = 0; i < len(ENODE); i++ {
		if string(ENODE[i]) != "@" {
			drawFH(10," ")
			fmt.Print(string(ENODE[i]))
		}
	}

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
			follow(&n[i], &n, i)
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
func follow(e *Edge, n *[8]Edge, x int) {
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
		myselect(&n[i], &n)
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
func myselect(e *Edge, n *[8]Edge) {
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

func first(e *Edge, n *[8]Edge, x int) {
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
		first(&n[i], &n, i)
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
