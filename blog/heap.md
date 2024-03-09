---
typora-root-url: ./
---

<a name="gDKN2"></a>
### container中定义的heap
在golang中的"container/heap"中定义了堆的实现，我们在使用时需要实现heap接口中定义的方法，以此实现一个堆。<br />在`container/heap.go`中的heap接口的定义如下：
```go
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}
```
而sort包中的接口定义如下：
```go
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```
所以我们实现一个堆时需要实现这五个方法，然后相当于实现了这个接口，然后就可以调用`container/heap.go`中定义的`Init`方法、`Push`方法、`Pop`方法进行堆的基础入堆、出堆操作。<br />在使用这三个方法时，需要注意按照源码中定义的函数的入参和返回值的类型来使用。
```go
// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}
```
```go
// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h Interface, x any) {
	h.Push(x)
	up(h, h.Len()-1)
}
```
```go
// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h Interface) any {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}
```

<a name="pg1kk"></a>
### heap的使用示例
在golang的源码中也有堆的使用示例：<br />可以看到实现上我们用切片来作为heap的底层实现类型。<br />下面的代码是定义一个小根堆的示例，如果我们想定义一个存int类型数据的大根堆，只需要把`Less`函数中的小于号换成大于号即可。
```go
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates an integer heap built using the heap interface.
package heap_test

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Example_intHeap() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

```
<a name="aZvJI"></a>
### 刷lc应用堆的示例
我们看一下[23. 合并 K 个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/)<br />

![](/image/lc.png)





这个题需要定义一个小根堆来存链表节点指针。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    h := minHeap{}
    for _, head := range lists {
        if head != nil {
           h = append(h, head) 
        }
    }     
    heap.Init(&h) 

    dummyhead := &ListNode{}
    cur := dummyhead
    
    for len(h)>0 {
        node := heap.Pop(&h).(*ListNode)
        if node.Next != nil {
            heap.Push(&h, node.Next)
        }
        cur.Next = node
        cur = cur.Next
    }
    return dummyhead.Next
}

type minHeap []*ListNode
func (h minHeap) Len() int {return len(h)}
func (h minHeap) Less(i,j int) bool {return h[i].Val<h[j].Val}
func (h minHeap) Swap(i,j int) { h[i], h[j] = h[j], h[i]}
func (h *minHeap) Push(x any) { *h = append(*h, x.(*ListNode))}
func (h *minHeap) Pop() any { old:=*h; n:=len(old); x:=old[n-1]; *h=old[:n-1]; return x}

```

