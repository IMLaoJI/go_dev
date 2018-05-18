package Sort

/**
* Created by LONG  on 2018/5/17.
*/
import (
	"go_dev/Algorithm/common"
	"math"
	"fmt"
)

const NBUCKET = 5
const INTERVAL = 10

type Node struct {
	data int
	next *Node
}

func BucketSortHigh(arr []int) {
	buckets := make([]*Node, NBUCKET)
	//for i := 0; i < NBUCKET; i++ {
	//	buckets = append(buckets, &Node{})

	//put items into the buckets
	for _, val := range arr {
		var current *Node
		pos := getBucketIndex(val)
		current = new(Node) //or current = &Node{}
		current.data = val
		current.next = buckets[pos]
		buckets[pos] = current
	}
	fmt.Println("Bucktets before sorted")
	//check what's in each bucket
	for i := 0; i < NBUCKET; i++ {
		fmt.Printf("Bucket[%d]: ", i)
		printBuckets(buckets[i])
		fmt.Println()
	}

	//sorting bucket using Insertion Sort
	for i := 0; i < NBUCKET; i++ {
		buckets[i] = InsertionSortByList(buckets[i])
	}
	fmt.Printf("\n\n")
	fmt.Println("Bucktets after sorted")
	//check what's in each bucket
	for i := 0; i < NBUCKET; i++ {
		fmt.Printf("Bucket[%d]: ", i)
		printBuckets(buckets[i])
		fmt.Println()
	}
	var j = 0
	// finally put items back to original array
	for i := 0; i < NBUCKET; i++ {
		var node *Node
		node = buckets[i]
		for node != nil {
			arr[j] = node.data
			j++
			node = node.next
		}
	}

}
func InsertionSortByList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	var nodelist = head
	var k = head.next
	nodelist.next = nil
	for k != nil {
		var ptr *Node
		if nodelist.data > k.data {
			var tmp *Node
			tmp = k
			k = k.next
			tmp.next = nodelist
			nodelist = tmp
			continue
		}
		for ptr = nodelist; ptr.next != nil; ptr = ptr.next {
			if ptr.next.data > k.data {
				break
			}
		}
		//中间有值大于k
		if ptr.next != nil {
			var tmp *Node
			tmp = k
			k = k.next
			tmp.next = ptr.next
			ptr.next = tmp
			continue
		} else { //没有值大于k 则插入头部 称为新头
			ptr.next = k
			k = k.next
			ptr.next.next = nil
			continue
		}
	}
	return nodelist

}
func printBuckets(head *Node) {
	cur := head
	for cur != nil {
		fmt.Printf(" %d", cur.data)
		cur = cur.next
	}
}
func getBucketIndex(val int) int {
	return val / INTERVAL
}

func main() {
	arr := common.GenerateRandomArray(20, 50)
	for index, val := range arr {
		arr[index] = int(math.Abs(float64(val)))
	}
	common.PrintArr(arr)
	BucketSortHigh(arr)
	fmt.Println(arr)
}
