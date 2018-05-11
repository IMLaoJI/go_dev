package queue

import (
	"github.com/cheekybits/genny/generic"
	"sync"
)

/**
* Created by LONG  on 2018/5/11.
*/

type Item generic.Type
type ItemQueue struct {
	items []Item
	lock  sync.RWMutex
}

func (q *ItemQueue) New() *ItemQueue {
	q.items = []Item{}
	return q
}

func (q *ItemQueue) Enqueue(t Item) {
	q.lock.Lock()
	q.items = append(q.items, t)
	q.lock.Unlock()
}

func (q *ItemQueue) Dequeue() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

func (q *ItemQueue) Front() *Item {
	q.lock.RLock()
	item := q.items[0]
	q.lock.RUnlock()
	return &item
}

// IsEmpty returns true if the queue is empty
func (q *ItemQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of Items in the queue
func (q *ItemQueue) Size() int {
	return len(q.items)
}