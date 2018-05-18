package Sort

/**
* Created by LONG  on 2018/5/18.
*/

import "go_dev/Algorithm/common"

func HeapSort(arr []int) {
	if arr == nil && len(arr) < 2 {
		return
	}
	for index := range arr {
		HeapInsert(arr, index)
	}

	var size = len(arr)
	size--
	common.Swap(arr, 0, size)

	for size > 0 {
		heapify(arr, 0, size)
		size--
		common.Swap(arr, 0, size)
	}
}

func HeapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		common.Swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func heapify(arr []int, index int, size int) {
	var left = index*2 + 1
	for left < size {
		var largest = common.If(left+1 < size && arr[left+1] > arr[left], left+1, left).(int)
		largest = common.If(arr[largest] > arr[index], largest, index).(int)
		if largest == index {
			break
		}
		common.Swap(arr, largest, index)
		index = largest
		left = index*2 + 1
	}
}
func main() {
	arr := common.GenerateRandomArray(10, 10)
	common.PrintArr(arr)
	HeapSort(arr)
	common.PrintArr(arr)
}
