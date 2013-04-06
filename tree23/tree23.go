package tree23

import ( 
	"fmt"
	"log"
)

// sedgewick 2-3 tree
type 𝞃 int

func (t 𝞃) Less(other 𝞃) bool {
	return t < other
}

type Tree23𝞃 struct {
	root *_nodeTree23𝞃
}

func NewTree23𝞃() Tree23𝞃 {
	// root as nil marks as empty tree
	return Tree23𝞃{}
}

func (t *Tree23𝞃) Insert(key 𝞃) {
	if t.root == nil {
		t.root = &_nodeTree23𝞃{}
		t.root.key1 = key
	} else {
		t.root.insert(key)
	}
}

type _nodeTree23𝞃 struct {
	key1, key2, key3 𝞃
	treetype int // 2tree, 3tree, 4tree
	lft, mid, rht, four *_nodeTree23𝞃
}

func (t *_nodeTree23𝞃) insert2(key 𝞃) {
	// only one key1 exists
	switch { 	
	case key.Less(t.key1):
		// key is less than key1
		switch {
		case t.lft == nil:
			// shift key1 to the right
			t.key1, t.key2 = key, t.key1
			t.treetype = 3
		case t.lft != nil:
			// don't touch this tree. yet.
			t.lft.insert(key)			
		}

	case t.key1.Less(key):
		switch {		
		case t.rht == nil:
			// key1 stays where it is
			// since the branches are nil, ignore the branches
			t.key2 = key
			t.treetype = 3
		case t.rht != nil:
			// don't touch this tree. yet.
			t.rht.insert(key)
		}

	case t.key1.Eq(key):
		



	}
	
	

}

func (t *_nodeTree23𝞃) insert(key 𝞃) {
	switch t.treetype {
	case 2: t.insert2(key)
	// case 3: t.insert3(key)
	// case 4: t.insert4(key)
	default:
		log.Panic("tree has invalid treetype of: " + fmt.Sprint(t.treetype))
	}
}







