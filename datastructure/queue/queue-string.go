// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package queue

import "sync"

/**
* Created by LONG  on 2018/5/11.
 */

type StringQueue struct {
	items []string
	lock  sync.RWMutex
}

func (q *StringQueue) New() *StringQueue {
	q.items = []string{}
	return q
}

func (q *StringQueue) Enqueue(t string) {
	q.lock.Lock()
	q.items = append(q.items, t)
	q.lock.Unlock()
}

func (q *StringQueue) Dequeue() *string {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

func (q *StringQueue) Front() *string {
	q.lock.RLock()
	item := q.items[0]
	q.lock.RUnlock()
	return &item
}

// IsEmpty returns true if the queue is empty
func (q *StringQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of Strings in the queue
func (q *StringQueue) Size() int {
	return len(q.items)
}
