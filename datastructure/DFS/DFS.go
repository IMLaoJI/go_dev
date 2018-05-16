package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

/**
* Created by LONG  on 2018/5/12.
*/
var maze [][]int
var book [][]int
var reader = bufio.NewReader(os.Stdin)
var outx, outy int
var MIN = 9999
var next [4][2]int = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var row, col int

func main() {
	var i, j int
	var inx, iny int
	arr := ReadInput()
	row = arr[0]
	col = arr[1]
	for i := 1; i <= row+10; i++ {
		tmp := make([]int, 0, col+10)
		for j := 1; j <= col+10; j++ {
			tmp = append(tmp, 0)
		}
		maze = append(maze, tmp)
		book = append(book, tmp)
	}
	for i = 1; i <= row; i++ {
		arr := ReadInput()
		for j = 1; j <= col; j++ {
			maze[i][j] = arr[j-1]
		}
	}
	//fmt.Println("end")
	io := ReadInput()
	inx = io[0]
	iny = io[1]
	outx = io[2]
	outy = io[3]
	book[inx][iny] = 1
	DFS(inx, inx, 0)
	fmt.Println("\n the shortest path length is %d", MIN)
}

func DFS(inx, iny int, step int) {
	var tx, ty, k int
	if inx == outx && iny == outy {
		if step < MIN {
			MIN = step
		}
		return
	}
	for k = 0; k < len(next); k++ {
		tx = inx + next[k][0]
		ty = iny + next[k][1]
		if tx < 1 || tx > row || ty < 1 || ty > col {
			continue
		}
		if maze[tx][ty] == 0 && book[tx][ty] == 0 {
				book[tx][ty] = 1
				DFS(tx, ty,  step+1)
				book[tx][ty] = 0
		}
	}
	return

}
func ReadInput() (arr []int) {
	s, err := reader.ReadString('\n')
	s = strings.Trim(s, " \n")
	if err != nil {
		fmt.Println("input error")
	}
	split := strings.Split(s, " ")
	arr = make([]int, 0, len(split))
	for i := 0; i < len(split); i++ {
		tem, err := strconv.Atoi(split[i])
		if err != nil {
			panic("input error")
		}
		arr = append(arr, tem)
	}
	return
}
