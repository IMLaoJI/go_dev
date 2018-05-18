package Sort

/**
* Created by LONG  on 2018/5/16.
*/
import "go_dev/Algorithm/common"

func SelectionSort(arr []int)  {
	if arr == nil && len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr) -1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			minIndex = common.If(arr[j] < arr[minIndex], j, minIndex).(int)
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	arr := common.GenerateRandomArray(10, 10)
	common.PrintArr(arr)
	SelectionSort(arr)
	common.PrintArr(arr)
}