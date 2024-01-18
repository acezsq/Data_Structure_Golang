package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 创建一个新的链表
	l := list.New()

	// 在链表尾部添加元素
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	// 在链表头部添加元素
	l.PushFront(0)

	// 遍历链表并打印元素值
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 0 1 2 3
	}

	// 获取链表长度
	fmt.Println("Length:", l.Len()) // 4

	// 删除指定元素
	elementToRemove := l.Front().Next()
	l.Remove(elementToRemove)

	// 在指定元素之前插入元素
	elementToInsertBefore := l.Back()
	l.InsertBefore(4, elementToInsertBefore)

	// 在指定元素之后插入元素
	elementToInsertAfter := l.Front()
	l.InsertAfter(5, elementToInsertAfter)

	// 遍历链表并打印元素值
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 0 5 2 4 3
	}

	// 清空链表
	l.Init()

	// 遍历链表并打印元素值
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) //
	}
}
