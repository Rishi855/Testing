package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	key   int
	value int
}

type MinHeap []*Node

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(*Node)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func main() {
	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, &Node{key: 12, value: 12})
	heap.Push(h, &Node{key: 10, value: 10})
	heap.Push(h, &Node{key: 14, value: 14})
	for h.Len()>0{
		node := heap.Pop(h).(*Node)
		fmt.Printf("key: %d, value: %d\n",node.key,node.value)
	}
}
