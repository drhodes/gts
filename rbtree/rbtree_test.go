package rbtree

import (
	"testing"
	"fmt"
)

func TestInsert(t *testing.T) {
	tree := NewRBTree()
	fmt.Println(tree)
	tree.put(𝞃(1))
	tree.put(𝞃(2))
	tree.put(𝞃(3))
	fmt.Println(tree)
	tree.root.delete(𝞃(2))
	fmt.Println(show(tree.root, 0))
}

