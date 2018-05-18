package Sort

import (
	"fmt"
	"strconv"
	"go_dev/Algorithm/common"
	"math"
)

/**
* Created by LONG  on 2018/5/16.
*/

func RadixSortByLSD(arr []int) []int {
	if arr == nil && len(arr) < 2 {
		fmt.Println("NO NEED TO SORT")
		return arr
	}
	maxl := MaxLenByString(arr)
	return RadixCoreByLSD(arr, 0, maxl)
}
func RadixCoreByLSD(arr []int, digit, maxl int) []int { //核心排序机制时间复杂度 O( d( r+n ) )
	if digit >= maxl {
		return arr //排序稳定
	}
	radix := 10
	count := make([]int, radix)
	bucket := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		//a := GetDigit(arr[i], digit)
		//fmt.Println(a)
		count[GetDigit(arr[i], digit)]++
	}
	for i := 1; i < radix; i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		d := GetDigit(arr[i], digit)
		bucket[count[d]-1] = arr[i]
		count[d]--
	}
	return RadixCoreByLSD(bucket, digit+1, maxl)
}

func GetDigitByLSD(x, d int) int { //获取某位上的数字
	a := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	return (x / a[d]) % 10
}

func MaxLenByStringByLSD(arr []int) int { //获取最大位数
	var maxl, curl int
	for i := 0; i < len(arr); i++ {
		curl = len(strconv.Itoa(arr[i]))
		if curl > maxl {
			maxl = curl
		}
	}
	return maxl
}

func MaxLenByMathByLSD(arr []int) int {
	var max = arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	var maxlen = 1
	for math.Floor(float64(max/10)) > 0 {
		maxlen++
		max = max / 10
	}
	return maxlen
}

func main() {
	arr := common.GenerateRandomArray(10, 1000)
	for index, val := range arr {
		arr[index] = int(math.Abs(float64(val)))
	}
	common.PrintArr(arr)
	num1 := MaxLenByString(arr)
	num2 := MaxLenByMath(arr)
	fmt.Println(num1,num2)
	arrtem := common.CopyArray(arr)
	common.Comparator(arrtem)
	common.PrintArr(arrtem)
	arr2 := RadixSort(arr)
	common.PrintArr(arr2)
	fmt.Println(common.IsEqual(arrtem, arr2))
}
