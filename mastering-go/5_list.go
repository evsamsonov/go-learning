package main

import (
	"container/list"
	"fmt"
)

func main() {
	doubleLinkedList := list.New()

	v1 := doubleLinkedList.PushBack("One")
	v2 := doubleLinkedList.PushBack("Two")
	doubleLinkedList.PushFront("Three")
	doubleLinkedList.InsertBefore("Four", v1)
	doubleLinkedList.InsertAfter("Five", v2)

	for dl := doubleLinkedList.Front(); dl != nil; dl = dl.Next() {
		fmt.Println(dl.Value)
	}
}
