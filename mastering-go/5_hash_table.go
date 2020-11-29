package main

import "fmt"

const hashTableSize = 15

type HashNode struct {
	Value int
	Next  *HashNode
}

type HashTable struct {
	Table map[int]*HashNode
	Size  int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		Table: make(map[int]*HashNode, size),
		Size:  size,
	}
}

func hashFunc(i, size int) int {
	return i % size
}

func (t *HashTable) Insert(value int) int {
	index := hashFunc(value, t.Size)
	node := HashNode{
		Value: value,
		Next:  t.Table[index],
	}
	t.Table[index] = &node
	return index
}

func (t *HashTable) Traverse() {
	for i := range t.Table {
		if t.Table[i] == nil {
			continue
		}
		node := t.Table[i]
		for node != nil {
			fmt.Printf("%d -> ", node.Value)
			node = node.Next
		}
		fmt.Println()
	}
}

func main() {
	hashTable := NewHashTable(hashTableSize)
	for i := 0; i < 120; i++ {
		hashTable.Insert(i)
	}
	hashTable.Traverse()
}
