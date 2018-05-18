package Sort

/**
* Created by LONG  on 2018/5/17.
*/
import (
	"go_dev/Algorithm/common"
	"math"
	"fmt"
)

//only for 0~200
func BucketSortLow(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	var max = -1
	for _, val := range arr {
		max = int(math.Max(float64(max), float64(val)))
	}
	var bucket = make([]int, max+1)
	for _, val := range arr {
		bucket[val]++
	}
	var i = 0
	for index, val := range bucket {
		for val >0 {
			arr[i] = index
			i++
			val--
		}
	}

}

func main() {
	arr := common.GenerateRandomArray(10, 1000)
	for index, val := range arr {
		arr[index] = int(math.Abs(float64(val)))
	}
	common.PrintArr(arr)
	BucketSortLow(arr)
	fmt.Println(arr)
}
