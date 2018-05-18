package Sort

import "go_dev/Algorithm/common"
import "math"
/**
* Created by LONG  on 2018/5/16.
*/

func CountingSort(arr []int) {
	var i, j, k, idx int
	var min, max int
	max = arr[0]
	min = max
	for i = 1; i < len(arr); i++ {
		min = common.If(arr[i] < min, arr[i], min).(int)
		max = common.If(arr[i] > max, arr[i], max).(int)
	}

	k = max - min + 1
	var tem = make([]int, k)
	for i = 0; i < len(tem); i++ {
		tem[i] = 0
	}
	for i = 0; i < len(arr); i++ {
		tem[arr[i]-min]++
	}

	for i = min; i <= max; i++ {
		for j = 0; j < tem[i-min]; j++ {
			arr[idx] = i
			idx++
		}
	}
	return
}

func main() {
	arr := common.GenerateRandomArray(10, 10)
	for index, val := range arr {
		arr[index] = int(math.Abs(float64(val)))
	}
	common.PrintArr(arr)
	CountingSort(arr)
	common.PrintArr(arr)
}
