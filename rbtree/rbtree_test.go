package rbtree

import (
	"testing"
	"fmt"
)

func TestInsert(t *testing.T) {
	tree := NewRBTree()
	fmt.Println(tree)
	tree.put(𝞃(1), 12)
	tree.put(𝞃(2), 2)
	tree.put(𝞃(3), 5)
	tree.put(𝞃(4), 4)
	tree.put(𝞃(5), 7)
	fmt.Println(tree)
	fmt.Println(show(tree.root, 0))
}

