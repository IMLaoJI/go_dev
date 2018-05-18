package Sort

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

/**
* Created by LONG  on 2018/5/11.
*/

func BubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("错了啊")
	}
	trimmedInput := strings.Trim(line, " \n")
	newline := strings.Split(trimmedInput, " ")
	var arr []int
	for _, val := range newline {
		i, e := strconv.Atoi(val)
		if e != nil {
			panic("输入有误")
		}
		arr = append(arr, i)
	}
	BubbleSort(arr)
	fmt.Print(arr)

}
