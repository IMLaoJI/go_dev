package Sort

/**
* Created by LONG  on 2018/5/18.
*/

import (
	"go_dev/Algorithm/common"
	"fmt"
)

func ShellSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	var h = 1
	size := len(arr)
	for h < size/3 {
		h = 3*h + 1
	}
	var i, j int
	for h >= 1 {
		for i = h; i < size; i++ {
			for j = i; j >= h && (arr[j] < arr[j-h]); j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		fmt.Printf("the step= %d   ", h)
		for k := 0; k < size; k++ {
			fmt.Printf("%d ", arr[k])
		}
		fmt.Println()
		h = h / 3
	}
}

func main() {
	arr := common.GenerateRandomArray(30, 10)
	common.PrintArr(arr)
	ShellSort(arr)
	common.PrintArr(arr)
}
