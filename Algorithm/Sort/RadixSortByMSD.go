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
type buc struct {
	num int
	arr []int
}

func RadixSort(arr []int) []int {
	if arr == nil && len(arr) < 2 {
		fmt.Println("NO NEED TO SORT")
		return arr
	}
	maxl := MaxLenByString(arr)
	RadixCore(arr, maxl)
	return arr
}
func RadixCore(arr []int, maxl int) { //核心排序机制时间复杂度 O( d( r+n ) )
	var radix = int(math.Pow(10, float64(maxl-1)))
	var bucket []*buc
	for i := 0; i < 10; i++ {
		bucket = append(bucket, &buc{num: 0, arr: []int{}})
	}
	for _, val := range arr {
		k := int(math.Floor(float64(val % int(math.Pow(10, float64(maxl))) / radix)))
		bucket[k].num++
		bucket[k].arr = append(bucket[k].arr, val)
	}
	//递归开始
	for j := 0; j < 10; j++ {
		if bucket[j].num > 1 {
			RadixCore(bucket[j].arr, maxl-1)
		}
	}
	var idx =0
	lengn := len(arr)
	//收集桶中数据 将其变有序
	for j := 0; j < 10; j++ {
		for _,val :=range bucket[j].arr {
			if idx < lengn {
				arr[idx] = val
				idx++
			}

		}
	}
	//return arrtem
}

func GetDigit(x, d int) int { //获取某位上的数字
	a := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	return (x / a[d]) % 10
}

func MaxLenByString(arr []int) int { //获取最大位数
	var maxl, curl int
	for i := 0; i < len(arr); i++ {
		curl = len(strconv.Itoa(arr[i]))
		if curl > maxl {
			maxl = curl
		}
	}
	return maxl
}

func MaxLenByMath(arr []int) int {
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
	fmt.Println(num1, num2)
	sort := RadixSort(arr)
	fmt.Println(sort)
	//arrtem := common.CopyArray(arr)
	//common.Comparator(arrtem)
	//common.PrintArr(arrtem)
	//arr2 := RadixSort(arr)
	//common.PrintArr(arr2)
	//fmt.Println(common.IsEqual(arrtem, arr2))
}
