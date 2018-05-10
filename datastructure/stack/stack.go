package stack

import (
	"sync"

)

// Item the type of the stack
//type Item rune

// ItemStack the stack of Items
type ItemStack struct {
	items []rune
	lock  sync.RWMutex
}

// New creates a new ItemStack
func (s *ItemStack) New() *ItemStack {
	s.items = []rune{}
	return s
}

// Push adds an Item to the top of the stack
func (s *ItemStack) Push(t rune) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Pop removes an Item from the top of the stack
func (s *ItemStack) Pop() *rune {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item
}

func (s *ItemStack)Peek() *rune {
	s.lock.RLock()
	defer s.lock.RUnlock()
	item:=*new(rune)
	if len(s.items) != 0 {
		item = s.items[len(s.items)-1]
		return &item
	}else {
		return &item
	}
	return &item
}

func (s *ItemStack) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *ItemStack) Size() int{
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}

type ItemStackInt struct {
	items []int
	lock  sync.RWMutex
}

// New creates a new ItemStack
func (s *ItemStackInt) New() *ItemStackInt {
	s.items = []int{}
	return s
}

// Push adds an Item to the top of the stack
func (s *ItemStackInt) Push(t int) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Pop removes an Item from the top of the stack
func (s *ItemStackInt) Pop() *int {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item
}

func (s *ItemStackInt)Peek() *int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	item:=*new(int)
	if len(s.items) != 0 {
		item = s.items[len(s.items)-1]
		return &item
	}else {
		return &item
	}
	return &item
}

func (s *ItemStackInt) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *ItemStackInt) Size() int{
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}