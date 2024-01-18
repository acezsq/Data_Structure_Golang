package main

import "fmt"

type stack struct {
	tt   int
	data []interface{}
}

func (stk *stack) push(k interface{}) {
	stk.data[stk.tt] = k
	stk.tt++
}

func (stk *stack) pop() {
	stk.tt--
}

func (stk *stack) len() int {
	return stk.tt
}

func (stk *stack) top() interface{} {
	k := stk.tt - 1
	return stk.data[k]
}

func (stk *stack) empty() bool {
	if stk.tt == 0 {
		return true
	} else {
		return false
	}
}

func NewStack() *stack {
	return &stack{
		tt:   0,
		data: make([]interface{}, 10000),
	}
}

func main() {
	stk1 := NewStack()
	stk1.push(5)
	stk1.push(10)
	stk1.push(15)
	stk1.pop()
	fmt.Println(stk1.top())
	fmt.Println(stk1.len())
	fmt.Println(stk1.empty())
}
