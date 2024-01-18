package main

import "fmt"

type queue struct {
	hh, tt int
	data   []interface{}
}

func (q *queue) push(k interface{}) {
	q.tt++
	q.data[q.tt] = k
}

func (q *queue) pop() interface{} {
	res := q.data[q.hh]
	q.hh++
	return res
}

func (q *queue) len() int {
	return q.tt - q.hh + 1
}

func (q *queue) empty() bool {
	if q.tt < q.hh {
		return true
	} else {
		return false
	}
}

func (q *queue) front() interface{} {
	return q.data[q.hh]
}

func (q *queue) back() interface{} {
	return q.data[q.tt]
}

func NewQueue() *queue {
	return &queue{
		hh:   0,
		tt:   -1,
		data: make([]interface{}, 10000),
	}
}

func main() {
	q1 := NewQueue()
	q1.push(5)
	q1.push(10)
	q1.push(15)
	q1.pop()
	fmt.Println(q1.empty())
	fmt.Println(q1.len())
	fmt.Println(q1.front())
	fmt.Println(q1.back())
	q1.pop()
	q1.pop()
	fmt.Println(q1.empty())
}
