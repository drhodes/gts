package main

// based on C implementation @
// http://www.eternallyconfuzzled.com/tuts/datastructures/jsw_tut_rbtree.aspx

// Red Black balanced tree library
// > Created (Julienne Walker): August 23, 2003
// > Modified (Julienne Walker): March 14, 2008
// > Ported to Go (Derek Rhodes): April 06, 2013

import (
	"fmt"
	"math/rand"
	"time"
)

type 𝞃 *int
type size_t int
const HEIGHT_LIMIT = 64 // Tallest allowable tree 

type jsw_rbnode struct {
	red int     // Color (1=red, 0=black) 
	data 𝞃;    // User-defined content 
	link [2]*jsw_rbnode // Left (0) and right (1) links 
}

type cmp_f func(𝞃, 𝞃) int
type dup_f func(𝞃) 𝞃
type rel_f func(𝞃)

type jsw_rbtree struct {
	root *jsw_rbnode; // Top of the tree 
	cmp         cmp_f;  // Compare two items 
	dup         dup_f;  // Clone an item (user-defined) 
	rel         rel_f;  // Destroy an item (user-defined) 
	size        size_t; // Number of items (user-defined) 
}

type jsw_rbtrav struct {
	tree *jsw_rbtree;               // Paired tree 
	it *jsw_rbnode;                 // Current node 
	path [HEIGHT_LIMIT]*jsw_rbnode; // Traversal path 
	top        size_t;                // Top of stack 
}

func main() {
	cmp := func(x, y 𝞃) int {
		switch {
		case *x < *y: return -1
		case *x == *y: return 0
		}			
		return 1		
	}
	dup := func(x 𝞃) 𝞃 {
		d := *x
		return &d
	}
	nop := func(x 𝞃){}

	for j := 0; j<10000; j+=1000 {
		t := jsw_rbnew(cmp, dup, nop)
		begin := time.Now()
		vals := []int{}
		for i:=0; i<j; i++ {
			vals = append(vals, rand.Intn(10))
		}
		for i:=0; i<j; i++ {
			jsw_rbinsert(t, 𝞃(&vals[i]))
		}

		for i:=0; i<j; i++ {
			jsw_rberase(t, 𝞃(&vals[i]))
		}
		end := time.Now()

		fmt.Println(end.Sub(begin).Seconds(), j)
		fmt.Println(t)
	}
}

func is_red(root *jsw_rbnode) bool {
  return root != nil && root.red == 1
}

func exclaim(x int) int {
	if x == 1 {
		return 0
	}
	return 1
}

func jsw_single(root *jsw_rbnode, dir int) *jsw_rbnode {
	save := root.link[exclaim(dir)]

	root.link[exclaim(dir)] = save.link[dir]
	save.link[dir] = root

	root.red = 1
	save.red = 0
	return save
}

func jsw_double ( root *jsw_rbnode, dir int ) *jsw_rbnode {
	root.link[exclaim(dir)] = jsw_single ( root.link[exclaim(dir)], exclaim(dir) )
  return jsw_single ( root, dir )
}

func new_node ( tree *jsw_rbtree, data 𝞃 ) *jsw_rbnode {
	rn := new(jsw_rbnode)

	if ( rn == nil ) {
		return nil
	}
	rn.red = 1
	rn.data = tree.dup ( data )
	rn.link[0] = nil
	rn.link[1] = nil

  return rn
}

func jsw_rbnew ( cmp cmp_f, dup dup_f, rel rel_f ) *jsw_rbtree {
	rt := new(jsw_rbtree)

	if ( rt == nil ) {
		return nil
	}

	rt.root = nil
	rt.cmp = cmp
	rt.dup = dup
	rt.rel = rel
	rt.size = 0
	
	return rt
}

func jsw_rbdelete ( tree *jsw_rbtree ) {
	it := tree.root
	var save *jsw_rbnode

    // Rotate away the left links so that
    // we can treat this like the destruction
    // of a linked list

	for it != nil {
		if it.link[0] == nil {
			// No left links, just kill the node and move on 
			save = it.link[1]
			tree.rel ( it.data )
			it = nil //free ( it )
		} else {
			// Rotate away the left link and check again 
			save = it.link[0]
			it.link[0] = save.link[1]
			save.link[1] = it
		}		
		it = save
	}

	//free ( tree )
	tree = nil

}


func LameB2I(b bool) int {
	if b {
		return 1
	}
	return 0
}

func jsw_rbfind ( tree *jsw_rbtree, data 𝞃) 𝞃 {
	it := tree.root

	for it != nil {
		cmp := tree.cmp ( it.data, data )

		if ( cmp == 0 ) {
			break
		}

		// If the tree supports duplicates, they should be
		// chained to the right subtree for this to work
		it = it.link[LameB2I(cmp < 0)]
	}
	
	if it == nil {
		return nil
	}
	return it.data
}

func jsw_rbinsert ( tree *jsw_rbtree, data 𝞃 ) int {
	if ( tree.root == nil ) {
		// We have an empty tree; attach the
		// new node directly to the root
		tree.root = new_node ( tree, data )
		
		if tree.root == nil {
			return 0
		}
	} else {
		head := jsw_rbnode{}  // False tree root 
		var g, t *jsw_rbnode  // Grandparent & parent 
		var p, q *jsw_rbnode;    // Iterator & parent 
		var dir, last int			

		// Set up our helpers 
		t = &head
		g = nil
		p = nil
		q = tree.root
		t.link[1] = q	
		
		// Search down the tree for a place to insert 
		for {
			if ( q == nil ) {
				// Insert a new node at the first null link 
				q = new_node ( tree, data )
				p.link[dir] = q

				if ( q == nil ) {
					return 0
				}
			} else if is_red( q.link[0] ) && is_red( q.link[1] ) {
				// Simple red violation: color flip 
				q.red = 1
				q.link[0].red = 0
				q.link[1].red = 0
			}
			
			if ( is_red ( q ) && is_red ( p ) ) {
				// Hard red violation: rotations necessary 

				// TODO:
				// port impedence, bool <-> int. should probably change the 
				// whole implementation to bool after this is working
				dir2 := 0
				if t.link[1] == g {
					dir2 = 1
				}
				
				if ( q == p.link[last] ) {
					t.link[dir2] = jsw_single ( g, exclaim(last))
				} else {
					t.link[dir2] = jsw_double ( g, exclaim(last))
				}
			}
			
			// Stop working if we inserted a node. This
			// check also disallows duplicates in the tree
			if tree.cmp( q.data, data) == 0 {
				break
			}

			last = dir
			dir = LameB2I(tree.cmp ( q.data, data ) < 0)

			// Move the helpers down 
			if ( g != nil ) {
				t = g
			}
			
			g, p = p, q
			q = q.link[dir]
		}
		
		// Update the root (it may be different) 
		tree.root = head.link[1]
	}

	// Make the root black for simplified logic 
	tree.root.red = 0
	tree.size++

	return 1
}

	
func jsw_rberase ( tree *jsw_rbtree, data 𝞃) int {
	if tree.root != nil {
		head := jsw_rbnode{} // False tree root 
		var q, p, g *jsw_rbnode // Helpers 
		f := new(jsw_rbnode)  // Found item 
		f = nil
		var dir int = 1
		
		// Set up our helpers 
		q = &head
		g, p = nil, nil
		q.link[1] = tree.root
		
		// Search and push a red node down
		// to fix red violations as we go
		for  q.link[dir] != nil  {
			last := dir
			
			// Move the helpers down 
			g, p = p, q
			q = q.link[dir]
			dir = LameB2I(tree.cmp ( q.data, data ) < 0)
			
			// Save the node with matching data and keep
			// going; we'll do removal tasks at the end
			
			if tree.cmp ( q.data, data ) == 0 {
				f = q
			}
			
			// Push the red node down with rotations and color flips 
			if !is_red ( q ) && !is_red ( q.link[dir] ) {
				if ( is_red ( q.link[exclaim(dir)] ) ) {
					p = jsw_single ( q, dir )
					p.link[last] = p
				} else if !is_red ( q.link[exclaim(dir)] ) {
					s := p.link[exclaim(last)] 
					if s != nil {
						if !is_red ( s.link[exclaim(last)] ) && !is_red ( s.link[last] ) {
							// Color flip 
							p.red = 0
							s.red = 1
							q.red = 1
						} else {
							dir2 := 0
							if g.link[1] == p {
								dir2 = 1
							}
							
							if is_red ( s.link[last] ) {
								g.link[dir2] = jsw_double ( p, last )
							} else if is_red ( s.link[exclaim(last)] ) {
								g.link[dir2] = jsw_single ( p, last )
							}
							
							// Ensure correct coloring 
							q.red = 1
							g.link[dir2].red = 1
							g.link[dir2].link[0].red = 0
							g.link[dir2].link[1].red = 0
						}
					}
				}
			}
		}
		
		// Replace and remove the saved node 
		if f != nil {
			tree.rel ( f.data )
			f.data = q.data

			// TODO on bool <-> int refactor, chop this down.
			if p.link[1] == q {
				if q.link[0] == nil {
					p.link[1] = q.link[1]
				} else {
					p.link[1] = q.link[0]
				}
			} else {
				if q.link[0] == nil {
					p.link[0] = q.link[1]
				} else {
					p.link[0] = q.link[0]
				}
			}
			q = nil
			//free ( q )
		}


		// Update the root (it may be different) 
		tree.root = head.link[1]

		// Make the root black for simplified logic 
		if ( tree.root != nil ) {
			tree.root.red = 0
		}

		tree.size--
	}

	return 1
}

func jsw_rbsize ( tree *jsw_rbtree ) size_t {
	return tree.size
}

func jsw_rbtnew () *jsw_rbtrav {
	return new(jsw_rbtrav)
}

func jsw_rbtdelete ( trav *jsw_rbtrav ) {
	trav = nil
	//free ( trav )
}

func start ( trav *jsw_rbtrav, tree *jsw_rbtree, dir int ) 𝞃 {
	trav.tree = tree
	trav.it = tree.root
	trav.top = 0

	// Save the path for later traversal 
	if trav.it != nil {
		for trav.it.link[dir] != nil {
			trav.top++
			trav.path[trav.top] = trav.it
			trav.it = trav.it.link[dir]
		}
	}

	if trav.it == nil {
		return nil
	} 
	return trav.it.data
}

func move( trav *jsw_rbtrav, dir int ) 𝞃 {
	if ( trav.it.link[dir] != nil ) {
		// Continue down this branch 
		trav.top++
		trav.path[trav.top] = trav.it
		trav.it = trav.it.link[dir]

		for trav.it.link[exclaim(dir)] != nil {
			trav.top++
			trav.path[trav.top] = trav.it
			trav.it = trav.it.link[exclaim(dir)]
		}
	} else {
		// Move to the next branch 
		var last *jsw_rbnode 
		
		for {
			if ( trav.top == 0 ) {
				trav.it = nil
				break
			}
			
			last = trav.it
			trav.top--
			trav.it = trav.path[trav.top]
			if last != trav.it.link[dir] {
				break
			}
		} 
	}
	
	if trav.it == nil {
		return nil
	}
	return trav.it.data
}

func jsw_rbtfirst ( trav *jsw_rbtrav, tree *jsw_rbtree ) 𝞃 {
  return start ( trav, tree, 0 ); // Min value 
}

func jsw_rbtlast ( trav *jsw_rbtrav, tree *jsw_rbtree ) 𝞃 {
  return start ( trav, tree, 1 ); // Max value 
}

func jsw_rbtnext ( trav *jsw_rbtrav ) 𝞃 {
  return move ( trav, 1 ); // Toward larger items 
}

func jsw_rbtprev ( trav *jsw_rbtrav ) 𝞃 {
  return move ( trav, 0 ); // Toward smaller items 
}


