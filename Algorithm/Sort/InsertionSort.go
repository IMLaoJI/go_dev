package Sort

import (
	"fmt"
	"go_dev/Algorithm/common"
)

/**
* Created by LONG  on 2018/5/11.
*/

func InsertionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}
func main() {
	fmt.Println("ss")
	arr := common.GenerateRandomArray(10, 10)
	common.PrintArr(arr)
	InsertionSort(arr)
	common.PrintArr(arr)
	//fmt.Println(rand.(1))
}
