package rbtree

import (
	"testing"
	"fmt"
)

func TestInsert(t *testing.T) {
	tree := NewRBTree()
	fmt.Println(tree)
	tree.put(𝞃(1))
	tree.put(𝞃(1))
	tree.put(𝞃(1))
	fmt.Println(tree)
	fmt.Println(show(tree.root, 0))
}

