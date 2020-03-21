package huffman

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value *Node
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value.Weight < pq[j].value.Weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[0]
	old[0] = nil
	*pq = old[1:n]

	return item.value
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value *Node) {
	item.value = value
	heap.Fix(pq, item.index)
}
