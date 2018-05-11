package linked_list

import (
	"github.com/cheekybits/genny/generic"
	"sync"
	"fmt"
)

/**
* Created by LONG  on 2018/5/11.
*/

type Item generic.Type

type Node struct {
	content Item
	next    *Node
}

type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (ll *ItemLinkedList) Append(t Item) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	node := Node{t, nil}
	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	ll.size++
}

func (ll *ItemLinkedList) Insert(i int, t Item) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return fmt.Errorf("Index out of bounds")
	}
	addNode := Node{t, nil}
	if i == 0 {
		addNode.next = ll.head
		ll.head = &addNode
		return nil
	}

	node := ll.head
	j := 0
	for j < i-2 {
		j++
		node = node.next
	}
	addNode.next = node.next
	node.next = &addNode
	ll.size++
	return nil
}

func (ll *ItemLinkedList) RemoveAt(i int) (*Item, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return nil, fmt.Errorf("Index out of bounds")
	}
	node := ll.head
	j := 0
	for j < i-1 {
		j++
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	ll.size--
	return &remove.content, nil
}

func (ll *ItemLinkedList) IndexOf(t Item) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	j := 0
	for {
		if node.content == t {
			return j
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		j++
	}
}

func (ll *ItemLinkedList) IsEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	if ll.head == nil {
		return true
	}
	return false
}


func (ll *ItemLinkedList) Size() int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	size := 1
	last := ll.head
	for {
		if last == nil || last.next == nil {
			break
		}
		last = last.next
		size++
	}
	return size
}

func (ll *ItemLinkedList) String() {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	j := 0
	for {
		if node == nil {
			break
		}
		j++
		fmt.Print(node.content)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}


func (ll *ItemLinkedList) Head() *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	return ll.head
}