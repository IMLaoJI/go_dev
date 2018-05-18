package common

import (
	"strings"
	"fmt"
	"strconv"
	"bufio"
	"os"
	"sort"
	"time"
	"math/rand"
)

/**
* Created by LONG  on 2018/5/15.
*/

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

//读取用户输入 并返回一个int数组
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

//奇技淫巧 通过异或来实现两个数交换数值 注意 两值如果相等 会导致出现0
func Swap(arr []int, i int, j int) {

	if i == j {
		return
	}
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
//标准库内置排序
func Comparator(arr []int) {
	sort.Ints(arr)
}

//根据参数 随机生成数组
func GenerateRandomArray(maxSize int, maxValue int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var arr = make([]int, r.Intn(maxSize))
	for i := 0; i < len(arr); i++ {
		arr[i] = r.Intn(maxValue+1) - r.Intn(maxValue)
	}
	return arr
}

//打印数组
func PrintArr(arr []int) {
	if arr == nil {
		return
	}
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

// 拷贝数组
func CopyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	var res = make([]int, len(arr))
	for index := range res {
		res[index] = arr[index]
	}
	return res
}
// 判断两个数组是否相等
func IsEqual(arr1, arr2 []int) bool {
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
		return false
	}

	if arr1 == nil && arr2 == nil {
		return false
	}

	for index := range arr1 {
		if arr1[index] != arr2[index] {
			return false
		}
	}

	return true
}

//模拟三元表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
