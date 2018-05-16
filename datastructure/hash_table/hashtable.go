package hash_table

import (
	"github.com/cheekybits/genny/generic"
	"sync"
	"fmt"
)

/**
* Created by LONG  on 2018/5/12.
*/

type Key generic.Type

type Value generic.Type

type ValueHashTable struct {
	items map[int]Value
	lock  sync.RWMutex
}

// the hash() private function uses the famous Horner's method
// to generate a hash of a string with O(n) complexity
func hash(k Key) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

func (ht *ValueHashTable) Put(k Key, v Value) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := hash(k)
	if ht.items == nil {
		ht.items = make(map[int]Value)
	}
	ht.items[i] = v
}

func (ht *ValueHashTable) Remove(k Key) bool {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	if ht.items == nil {
		return false
	}
	i := hash(k)
	delete(ht.items, i)
	return true
}

func (ht *ValueHashTable) Get(k Key) Value {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	i := hash(k)
	if ht.items == nil {
		return nil
	}
	return ht.items[i]

}

func (ht *ValueHashTable) Sizes() int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return len(ht.items)
}