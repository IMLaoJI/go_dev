package main

/**
* Created by LONG  on 2018/5/18.
*/
import (
	"go_dev/Algorithm/common"
	"fmt"
)

/*
荷兰国旗问题
给定一个数组arr，和一个数num，请把小于num的数放在数组的
左边，等于num的数放在数组的中间，大于num的数放在数组的
右边。
*/
func Partition(arr []int, l, r, p int) []int {
	less := l - 1
	more := r + 1
	cur := l
	for cur < more {
		if arr[cur] < p {
			less++
			common.Swap(arr, less, cur)
			cur++
		} else if arr[cur] > p {
			more--
			common.Swap(arr, more, cur)
		} else {
			cur++
		}
	}
	return []int{less + 1, more - 1}
}

func main() {
	arr := common.GenerateRandomArray(30, 10)
	common.PrintArr(arr)
	res := Partition(arr, 0, len(arr)-1, 5)
	common.PrintArr(arr)
	fmt.Println(res[0])
	fmt.Println(res[1])
}
