package tree23

import ( 
	"fmt"
	"log"
)

// sedgewick 2-3 tree
// start dummy
type 𝞃 int

func (t 𝞃) Less(other 𝞃) bool {
	return t <= other
}
func (t 𝞃) Eq(other 𝞃) bool {
	return t == other
}
func New𝞃() 𝞃 {
	return 𝞃(0)
}
// end dummy

type nodetype int
const (
	type2 nodetype  = iota
	type3
	type4
)

type node struct {
	ntype nodetype
	k1, k2, k3 𝞃
	n1, n2, n3, n4 *node
}

func node2(key1 𝞃, n1, n2 *node) *node {
	return &node{type2, key1, New𝞃(), New𝞃(), n1, n2, nil, nil}
}
func node3(key1, key2 𝞃, n1, n2, n3 *node) *node {
	return &node{type3, key1, key2, New𝞃(), n1, n2, n3, nil}
}
func node4(key1, key2, key3 𝞃, n1, n2, n3, n4 *node) *node {
	return &node{type4, key1, key2, key3, n1, n2, n3, n4}
}

func (n *node) insert(key 𝞃) *node {	
	r := new(node)
	switch n.ntype {
	case type2: r = n.insert2(key)
	case type3: r = n.insert3(key)
	}
	if r.ntype == type4 {
		log.Println("type4 detected, splitting")
		//    [a] [b] [c]            [b]
		//   /   |   |   \    ->    /   \
		//  1    2   3    4      [a]     [c]
		//                      1   2   3   4
		a := node2(r.k1, r.n1, r.n2)
		c := node2(r.k3, r.n3, r.n4)
		b := node2(r.k2, a, c)
		r = b
	}
	return r
}
 
func (n *node) insert3(key 𝞃) *node {
	log.Println("Inserting3, ", key)
	switch {
	case key.Less(n.k1): 
		//    [b] [c]     ->     [+] [b] [c]
		//   /   |   \          /   |   |   \
		//  1    2    3        -    1   2    3
		if n.n1 == nil {
			return node4(key, n.k1, n.k2, nil, n.n1, n.n2, n.n3)
		} else {
			n.n1 = n.n1.insert(key)
			return n
		}

	case key.Eq(n.k1) || key.Less(n.k2) || key.Eq(n.k2):
		//    [a] [c]     ->     [a] [+] [c]
		//   /   |   \          /   |   |   \
		//  1    2    3        1    2   -    3
		if n.n2 == nil {
			return node4(n.k1, key, n.k2, n.n1, n.n2, nil, n.n3)
		} else {
			n.n2 = n.n2.insert(key)
			return n
		}

	case n.k3.Less(key):
		//    [a] [b]     ->     [a] [b] [+]
		//   /   |   \          /   |   |   \
		//  1    2    3        1    2   3    -
		if n.n3 == nil {
			return node4(n.k1, n.k2, key, n.n1, n.n2, n.n3, nil)
		} else {
			n.n3 = n.n3.insert(key)
			return n
		}
	} 
	log.Panic("hit deadpath")
	return nil 
}

func (n *node) insert2(key 𝞃) *node {
	log.Println("Inserting2, ", key)
	if key.Less(n.k1) {
		//    [a]     ->    [+] [a]
		//   /   \         /   |   \
		//  1     2       -    1    2   
		if n.n1 == nil {
			return node3(key, n.k1, nil, n.n1, n.n2)
		} else {
			n.n1 = n.n1.insert(key)
			return n
		}
	} else {
		//    [a]     ->    [a] [+]
		//   /   \         /   |   \
		//  1     2       1    2    -   
		if n.n2 == nil {
			return node3(n.k1, key, n.n1, n.n2, nil)
		} else {
			n.n2 = n.n2.insert(key)
			return n
		}
	}
	log.Panic("hit deadcode")
	return nil // deadcode
}

type Tree23𝞃 struct {
	root *node
}

func show(n *node, depth int) {
	padding := ""
	for i:=0; i<depth*4; i++ {
		padding += " "
	}
	if n == nil {
		//fmt.Println(padding, "nil")
		return 
	}
	fmt.Println(padding,n.k1, n.k2, n.k3)
	show(n.n1, depth+1)
	show(n.n2, depth+1)
	show(n.n3, depth+1)
	show(n.n4, depth+1)

}

func (t *Tree23𝞃) Show() {
	show(t.root, 0)
	
}

func (t *Tree23𝞃) Insert (key 𝞃) {
	fmt.Println(key)
	if t.root == nil {
		t.root = node2(key, nil, nil)
	} else {
		t.root = t.root.insert(key)
	}
}





