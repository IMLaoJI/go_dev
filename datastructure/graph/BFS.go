package graph

/**
* Created by LONG  on 2018/5/12.
*/

func (g *StringGraph) DFS(f func(*Node)) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	itemQueue := NodeQueue{}
	q := itemQueue.New()
	n := g.nodes[0]
	q.Enqueue(*n)
	visited := make(map[*Node]bool)
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		visited[node] = true
		near := g.edges[*node]
		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				q.Enqueue(*j)
				visited[j] = true
			}
		}
		if f != nil {
			f(node)
		}

	}

}
