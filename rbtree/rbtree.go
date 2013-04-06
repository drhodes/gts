package rbtree

// sedgewick 2007
// https://class.coursera.org/algs4partI-002/lecture/50

// Red links lean left
// No node has two red links connected to it
// Every path from root to null link has the same number of black links
// find is exactly the same as normal BST

import (
	"fmt"
)

// start dummy 
type 𝞃 int
type Val int

func (t 𝞃) CompareTo(other 𝞃) int {
	switch {
	case t < other: return -1
	case t > other: return 1
	case t == other: return 0
	}
	// dead path
	return -2
}

// end dummy

const (
	RED = true
	BLACK = false
)

type _node struct {
	key 𝞃
	val Val
	left, right *_node
	N int
	color bool
}

func new_node(key 𝞃, val Val, N int, color bool) *_node {
	n := new(_node)
	n.key = key
	n.val = val
	n.N = N
	n.color = color
	return n
}

// not recursive, size is calculated more efficiently.
func size(h *_node) int {
	if h == nil {
		return 0
	} else {
		return h.N
	}
}

func rotateLeft(h *_node) *_node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	return x
}

func isRed(x *_node) bool {	
	if x == nil {
		return false 
	}
	return x.color == RED
}

func rotateRight(h *_node) *_node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	x.N = h.N
	h.N = 1 + size(h.left) + 1 + size(h.right)	
	return x
}

func flipColors(h *_node) {
	h.color = RED;
	h.left.color = BLACK
	h.right.color = BLACK
}



type rbtree struct {
	root *_node
}

func NewRBTree() rbtree {
	return rbtree{}
}

func (t *rbtree) put(key 𝞃, val Val) {
	t.root = t.put2(t.root, key, val) 
	t.root.color = BLACK
}

func (t *rbtree) put2(h *_node, key 𝞃, val Val) *_node {
	if h == nil {
		return new_node(key, val, 1, RED)
	}

	cmp := key.CompareTo(h.key)
	switch {
	case cmp < 0:
		h.left = t.put2(h.left, key, val)
	case cmp > 0:
		h.right = t.put2(h.right, key, val)
	default:
		h.val = val
	}
	
	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}
	
	h.N = size(h.left) + size(h.right) + 1
	return h
}

func show(h *_node, depth int) string {
	padding := ""
	for i:=0; i<depth*4; i++ {
		padding += " "
	}
	if h == nil {
		return padding + "nil"
	}
	
	left := show(h.left, depth + 1) + "\n"
	right := show(h.right, depth + 1) + "\n"
	return fmt.Sprintf("%s%d: %d\n", padding, h.key, h.val) + left + right

}
