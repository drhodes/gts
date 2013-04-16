package rbtree

// ported version of 
// http://www.cs.princeton.edu/~rs/talks/LLRB/08Dagstuhl/Java/RedBlackBST.java

import (
	"log"
)

type (
	β int	
	α int
)

func Newβ() β {
	return β(0)
}

func (k α) compareTo(other α) int {
	if k < other {
		return -1
	} else if k > other {
		return 1 
	}
	return 0
}

const (
	RED = true
	BLACK = false
)

type node struct {
	key α
	value β
	left, right *node
	xc, yc float64
	color bool;      // color of parent link
}

func newNode(key α, value β, color bool) *node {
	n := node{}	
	n.key   = key
	n.value = value
	n.color = color
	return &n
}

type RedBlackBST struct {
	// all three fields are private, so change the names once tests pass
	root *node;      // root of the BST
	N int;           // size of BST
	k int
}

func get(x *node, key α) (β, bool) {
	if (x == nil) {
        return Newβ(), false
	}
	if eq(key, x.key) {
		return x.value, true
	}
	if less(key, x.key) {
		return get(x.left,  key)
	} else {
		return get(x.right, key)
    }
}

func (rb *RedBlackBST) get(key α) (β, bool) {  
	return get(rb.root, key)  
}

func (rb *RedBlackBST) contains(key α) bool {  
	_, b := rb.get(key)
	return b
}

func (rb *RedBlackBST) put(key α, value β) {
	if (!rb.contains(key)) {
		rb.N++
	}
	rb.root = insert(rb.root, key, value)
	rb.root.color = BLACK
}

func insert(h *node, key α, value β) *node {
	if (h == nil) {
		return newNode(key, value, RED)
	}
	if (isRed(h.left)) {
		if (isRed(h.left.left)) {
            h = splitFourNode(h)
		}
	}
	if (eq(key, h.key)) {
		h.value = value
	} else if (less(key, h.key)) {
		h.left = insert(h.left, key, value); 
	} else {
		h.right = insert(h.right, key, value); 
	} 

	if (isRed(h.right)) {
		h = leanLeft(h)
	} 

	return h
}

func (rb *RedBlackBST) deleteMin() {	
	rb.root = deleteMin(rb.root)
	rb.root.color = BLACK
	rb.N--
}

func deleteMin(h *node) *node { 	
	if (h.left == nil) {
		return nil
	}
	if (!isRed(h.left) && !isRed(h.left.left)) {
		h = moveRedLeft(h)
	}

	h.left = deleteMin(h.left)
	
	if (isRed(h.right)) {
		h = leanLeft(h)
	}

	return h
}

func (rb *RedBlackBST) deleteMax() { 
	rb.root = deleteMax(rb.root)
	rb.root.color = BLACK
	rb.N--
}

func deleteMax(h *node) *node { 
	if (h.right == nil) {  
		if (h.left != nil) {
            h.left.color = BLACK
			return h.left
		}		
	}

	if (isRed(h.left)) {
		h = leanRight(h)
	}
	
	if (!isRed(h.right) && !isRed(h.right.left)) {
		h = moveRedRight(h)
	}

	h.right = deleteMax(h.right)

	if (isRed(h.right)) {
		h = leanLeft(h)
	}

	return h
}

func (rb *RedBlackBST) Delete(key α) {
	if rb.N == 0 {
		log.Panic("Tried to delete an element from an empty αβRedBlackBST")
	}

	rb.root = delete(rb.root, key)
	rb.root.color = BLACK
	rb.N--
}

func delete(h *node, key α) *node { 
	if (less(key, h.key)) {
		if (!isRed(h.left) && !isRed(h.left.left)) {
            h = moveRedLeft(h)
			h.left =  delete(h.left, key)
		}
	} else {
		if (isRed(h.left)) {
			h = leanRight(h)
		}
		if (eq(key, h.key) && (h.right == nil)) {
			return nil
		}
		if (!isRed(h.right) && !isRed(h.right.left)) {
			h = moveRedRight(h)
		}
		if eq(key, h.key) {
			// 
			v, ok := get(h.right, min(h.right))
			if ok {
				h.value = v
			} else {
				h.value = Newβ()
			}
			// end
            h.key = min(h.right)
			h.right = deleteMin(h.right)
		} else {
			h.right = delete(h.right, key)
		}
	}
	
	if (isRed(h.right)) {
		h = leanLeft(h)
	} 
	return h
}

// Helper methods
func less(a, b α) bool { 
	return a.compareTo(b) <  0; 
}  

func eq  (a, b α) bool { 
	return a.compareTo(b) == 0; 
} 

func isRed(x *node) bool {
	if (x == nil) {
		return false
	}
	return (x.color == RED)
}

func rotR(h *node) *node {  // Rotate right.
	x := h.left
	h.left = x.right
	x.right = h
	return x
}

func rotL(h *node) *node {  // Rotate left.
	x := h.right
	h.right = x.left
	x.left = h
	return x
}

func splitFourNode(h *node)  *node {  // Rotate right, then flip colors.
	h = rotR(h)
	//      h.color       = RED
	h.left.color  = BLACK
	return h
}

func leanLeft(h *node)  *node {  // Make a right-leaning 3-node lean to the left.
	h = rotL(h)
	h.color      = h.left.color;                   
	h.left.color = RED;                     
	return h
}

func leanRight(h *node)  *node {  // Make a left-leaning 3-node lean to the right.
	h = rotR(h)
	h.color       = h.right.color;                   
	h.right.color = RED;                     
	return h
}

func moveRedLeft(h *node)  *node {  // Assuming that h is red and both h.left and h.left.left
	// are black, make h.left or one of its children red.
	h.color      = BLACK
	h.left.color = RED;  

	if (!isRed(h.right.left))  {
		h.right.color = RED; 
	} else { 
		h.right = rotR(h.right)
		h = rotL(h)
	}
	return h
}

func moveRedRight(h *node)  *node {  // Assuming that h is red and both h.right and h.right.left
	// are black, make h.right or one of its children red.
	h.color      = BLACK
	h.right.color = RED;  
	if (!isRed(h.left.left)) {
		h.left.color = RED; 
	} else { 
		h = rotR(h)
		h.color = RED
		h.left.color = BLACK
	}
	return h
}

// Utility functions

// return number of key-value pairs in symbol table
func (rb *RedBlackBST) size() int { 
	return rb.N
}  

// height of tree (empty tree height = 0)
func (rb *RedBlackBST) height() int { 
	return height(rb.root) 
}  

func MaxInt(a, b int) int {
	if a > b {
		return a 
	}
	return b
}

func height(x *node) int {  
	if (x == nil) {
		return 0
	}
	return 1 + MaxInt(height(x.left), height(x.right))
}

// return the smallest key
func (rb *RedBlackBST) min() α {
	return min(rb.root)
}

func min(x *node) α {
	var key α
	for ; x != nil; x = x.left {
		key = x.key
	}
	return key
}

// return the largest key
func (rb *RedBlackBST) max() α {
	var key α
	for x := rb.root; x != nil; x = x.right {
		key = x.key
	}
	return key
}



// Integrity checks

func (rb *RedBlackBST) check()   bool {  // Is this tree a red-black tree?
	return rb.isBST() && rb.is234() && rb.isBalanced()
}

func (rb *RedBlackBST) isBST() bool {  // Is this tree a BST?
	return isBST(rb.root, rb.min(), rb.max())
}

// Are all the values in the BST rb.rooted at x between min and max,
// and does the same property hold for both subtrees?
func isBST(x *node, min, max α) bool {  
	if (x == nil) {
		return true
	}
	if less(x.key, min) || less(max, x.key) {
		return false
	}
	return isBST(x.left, min, x.key) && isBST(x.right, x.key, max)
} 

func (rb *RedBlackBST) is234() bool { 
	return is234(rb.root); 
}  

// Does the tree have no red right links, and at most two (left)
// red links in a row on any path?
func is234(x *node)  bool { 
	if (x == nil) {
		return true
	}
	if (isRed(x.right)) {
		return false
	}
	if (isRed(x)) {
		if (isRed(x.left)) {
			if (isRed(x.left.left)) {
				return false
			}
		}
	}
	return is234(x.left) && is234(x.right)
} 

// Do all paths from rb.root to leaf have same number of black edges?
// number of black links on path from rb.root to min
func (rb *RedBlackBST) isBalanced()  bool { 
	black := 0
	x := rb.root
	for x != nil {
		if !isRed(x) {
			black++
		}
		x = x.left
	}
	return isBalanced(rb.root, black)
}

// Does every path from the rb.root to a leaf have the given number 
func isBalanced(x *node, black int) bool { 
	// of black links?
	if x == nil && black == 0 {
		return true
	} else if (x == nil && black != 0) {
		return false
	}
	if (!isRed(x)) {
		black--
	}
	return isBalanced(x.left, black) && isBalanced(x.right, black)
} 

