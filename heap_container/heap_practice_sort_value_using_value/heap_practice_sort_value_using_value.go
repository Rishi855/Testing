package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	key   int
	value int
}

type PairHeap []Pair

func (h PairHeap) Len() int { return len(h) }

func (h PairHeap) Less(i, j int) bool { return h[i].value < h[j].value }

func (h PairHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x any) { *h = append(*h, x.(Pair)) }

func (h *PairHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func main() {
	h := &PairHeap{}
	heap.Init(h)

	heap.Push(h, Pair{10, 3})
	heap.Push(h, Pair{2, 9})
	heap.Push(h, Pair{7, 15})
	heap.Push(h, Pair{8, 13})
	heap.Push(h, Pair{5, 15})

	for h.Len() > 0 {
		p := heap.Pop(h).(Pair)
		fmt.Printf("key: %d, value: %d\n", p.key, p.value)
	}
}
