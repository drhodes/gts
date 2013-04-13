package rbtree

// implementation based on:
// sedgewick 2007
// algorithms 4th edition
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

func (t 𝞃) CompareTo(other 𝞃) int {
	if t < other {
		return -1
	} else if t > other {
		return 1
	}
	// t == other.
	return 0
}

func 

// end dummy




const (
	RED = true
	BLACK = false
)

type _node struct {
	key 𝞃
	left, right *_node
	N int
	color bool
}

func new_node(key 𝞃, N int, color bool) *_node {
	n := new(_node)
	n.key = key
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

func (t rbtree) Show() string {
	return show(t.root, 0)
}


func (t *rbtree) put(key 𝞃) {
	t.root = t.put2(t.root, key)
	t.root.color = BLACK
}

func (t *rbtree) put2(h *_node, key 𝞃) *_node {
	if h == nil {
		return new_node(key, 1, RED)
	}

	cmp := key.CompareTo(h.key)
	switch {
	case cmp <= 0:
		h.left = t.put2(h.left, key)
	case cmp > 0:
		h.right = t.put2(h.right, key)
	default:
		//h.val = val
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
	return fmt.Sprintf("%s%d\n", padding, h.key) + left + right
}

func colorFlip(h *_node) *_node { 
	h.color = !h.color;	
	h.left.color = !h.left.color;
	h.right.color = !h.right.color;	
	return h;
}

func min(h *_node) 𝞃 {
	// implement this
	return New𝞃()
}

func moveRedLeft(h *_node) *_node { 
	colorFlip(h)
	
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		colorFlip(h)
	}
	return h
}

// Assuming that h is red and both h.right and h.right.left
// are black, make h.right or one of its children red.
func moveRedRight(h *_node) *_node { 
	h.color = BLACK
	h.right.color = RED
	if !isRed(h.left.left) {
		h.left.color = RED
	} else { 
		h = rotateRight(h)
		h.color = RED
		h.left.color = BLACK
	}
	return h
}

// Make a left-leaning 3-node lean to the right.
func leanRight(h *_node) *_node {  
	h = rotateRight(h);
	h.color = h.right.color;                   
	h.right.color = RED;                     
	return h
}

func delete(h *_node, key 𝞃) *_node {
	cmp := key.CompareTo(h.key)

	if cmp < 0 {
		if (!isRed(h.left) && !isRed(h.left.left)) {			
			h = moveRedLeft(h);			
			h.left = delete(h.left, key);			
		}
	} else {		
		if (isRed(h.left)) {
			h = leanRight(h);
		}

		if (cmp == 0 && (h.right == nil)) {
			return nil;
		}
		
		if !isRed(h.right) && !isRed(h.right.left) {
			h = moveRedRight(h)
		}
		
		if (cmp == 0) {			
			h.key = min(h.right);			
			h.value = get(h.right, h.key);			
			h.right = deleteMin(h.right);			
		} else {
			h.right = delete(h.right, key);
		}
	}
	return fixUp(h);
}
