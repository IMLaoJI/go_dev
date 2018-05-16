package graph

//import (
//	"github.com/cheekybits/genny/generic"
//	"fmt"
//	"sync"
//)
//
///**
//* Created by LONG  on 2018/5/12.
//*/
//
//type NodeQueue struct {
//	items []Node
//	lock  sync.RWMutex
//}
//
//// New creates a new NodeQueue
//func (s *NodeQueue) New() *NodeQueue {
//	s.lock.Lock()
//	s.items = []Node{}
//	s.lock.Unlock()
//	return s
//}
//
//// Enqueue adds an Node to the end of the queue
//func (s *NodeQueue) Enqueue(t Node) {
//	s.lock.Lock()
//	s.items = append(s.items, t)
//	s.lock.Unlock()
//}
//
//// Dequeue removes an Node from the start of the queue
//func (s *NodeQueue) Dequeue() *Node {
//	s.lock.Lock()
//	item := s.items[0]
//	s.items = s.items[1:len(s.items)]
//	s.lock.Unlock()
//	return &item
//}
//
//// Front returns the item next in the queue, without removing it
//func (s *NodeQueue) Front() *Node {
//	s.lock.RLock()
//	item := s.items[0]
//	s.lock.RUnlock()
//	return &item
//}
//
//// IsEmpty returns true if the queue is empty
//func (s *NodeQueue) IsEmpty() bool {
//	s.lock.RLock()
//	defer s.lock.RUnlock()
//	return len(s.items) == 0
//}
//
//// Size returns the number of Nodes in the queue
//func (s *NodeQueue) Size() int {
//	s.lock.RLock()
//	defer s.lock.RUnlock()
//	return len(s.items)
//}
//
//type Item generic.Type
//
//type Node struct {
//	value Item
//}
//
//func (n *Node) String() string {
//	return fmt.Sprintf("%v", n.value)
//}
//
//type ItemGraph struct {
//	nodes []*Node
//	edges map[Node][]*Node
//	lock  sync.RWMutex
//}
//
//func (g *ItemGraph) AddNode(n *Node) {
//	g.lock.Lock()
//	defer g.lock.Unlock()
//	g.nodes = append(g.nodes, n)
//}
//
//func (g *ItemGraph) AddEdge(n1, n2 *Node) {
//	g.lock.Lock()
//	defer g.lock.Unlock()
//	if g.edges == nil {
//		g.edges = make(map[Node][]*Node)
//	}
//	g.edges[*n1] = append(g.edges[*n1], n2)
//	g.edges[*n2] = append(g.edges[*n2], n1)
//}
//
//func (g *ItemGraph) String() {
//	g.lock.RLock()
//	defer g.lock.RUnlock()
//	s := ""
//	for i := 0; i < len(g.nodes); i++ {
//		s += g.nodes[i].String() + " -> "
//		near := g.edges[*g.nodes[i]]
//		for j := 0; j < len(near); j++ {
//			s += near[j].String() + " "
//		}
//		s += "\n"
//	}
//	fmt.Println(s)
//}
//
//func (g *ItemGraph) Traverse(f func(*Node)) {
//	g.lock.RLock()
//	defer g.lock.RUnlock()
//	itemQueue := NodeQueue{}
//	q := itemQueue.New()
//	n := g.nodes[0]
//	q.Enqueue(*n)
//	visited := make(map[*Node]bool)
//	for {
//		if q.IsEmpty() {
//			break
//		}
//		node := q.Dequeue()
//		visited[node] = true
//		near := g.edges[*node]
//		for i := 0; i < len(near); i++ {
//			j := near[i]
//			if !visited[j] {
//				q.Enqueue(*j)
//				visited[j] = true
//			}
//		}
//		if f != nil {
//			f(node)
//		}
//
//	}
//
//}


