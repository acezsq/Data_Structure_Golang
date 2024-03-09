package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j] // 大根堆
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &MaxHeap{}
	heap.Init(h)

	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 1)

	fmt.Println("当前最大值：", (*h)[0]) // 输出：3

	heap.Push(h, 5)
	heap.Push(h, 4)

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
