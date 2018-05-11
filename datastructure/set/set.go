package set

import "github.com/cheekybits/genny/generic"

/**
* Created by LONG  on 2018/5/11.
*/
type Item generic.Type
type ItemSet struct {
	items map[Item]bool
}

func (s *ItemSet) Add(item Item) *ItemSet {
	if s.items == nil {
		s.items = make(map[Item]bool)
	}
	_, ok := s.items[item]
	if !ok {
		s.items[item] = true
	}
	return s
}

func (s *ItemSet) Delete(item Item) bool {
	if s.items == nil {
		s.items = make(map[Item]bool)
	}
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

func (s *ItemSet) Clear() {
	s.items = make(map[Item]bool)
}

func (s *ItemSet) Size() int {
	return len(s.items)
}

func (s *ItemSet) Has(item Item) bool {
	_, ok := s.items[item]
	return ok
}
func (s *ItemSet) Items() []Item {
	items := []Item{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}
