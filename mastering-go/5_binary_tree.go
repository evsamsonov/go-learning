package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNotFound = errors.New("not found")

type Tree struct {
	Root *Node
}

type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
}

func main() {
	tree := NewTree(1, "Ivan")
	tree.Insert(2, "Dmitry")
	tree.Insert(-3, "Vladimir")
	tree.Insert(10, "Darya")

	name, err := tree.Search(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	name, err = tree.Search(77)
	if err != nil {
		fmt.Println("Not found!")
	}

	tree.Walk()
}

func NewTree(rootKey int, rootVal string) *Tree {
	return &Tree{
		Root: &Node{
			Key:   rootKey,
			Value: rootVal,
		},
	}
}

func (t *Tree) Insert(key int, val string) {
	if t.Root == nil {
		t.Root = &Node{
			Key:   key,
			Value: val,
		}
		return
	}
	t.doInsert(t.Root, key, val)
}

func (t *Tree) doInsert(node *Node, key int, val string) {
	if key == node.Key {
		node.Value = val
		return
	}
	var n **Node
	if key > node.Key {
		n = &node.Right
	} else {
		n = &node.Left
	}
	if *n == nil {
		*n = &Node{
			Key:   key,
			Value: val,
		}
		return
	}
	t.doInsert(*n, key, val)
}

func (t *Tree) Search(key int) (string, error) {
	return t.doSearch(t.Root, key)
}

func (t *Tree) doSearch(node *Node, key int) (string, error) {
	if node == nil {
		return "", ErrNotFound
	}
	if key == node.Key {
		return node.Value, nil
	}
	if key > node.Key {
		return t.doSearch(node.Right, key)
	}
	return t.doSearch(node.Left, key)
}

func (t *Tree) Walk() {
	t.doWalk(t.Root)
}

func (t *Tree) doWalk(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d: %s\n", node.Key, node.Value)
	t.doWalk(node.Left)
	t.doWalk(node.Right)
}
