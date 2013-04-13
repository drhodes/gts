package main

// jelly tree.  
// inspired by a 2-3 tree, but easier to implement.
// Derek A. Rhodes 4/9/13

import (
	"fmt"
	//"errors"
)

// ------------------------------------------------------------------
type node struct {
	Group
	left, right *node
}

func newNode(k int) *node {
	g := NewGroup(k)
	return &node{g, nil, nil}
}

func (n *node) split() error {
	fmt.Println("'''''''''''''''''''''''''''''''''")
	fmt.Println(n.Group.Show())
	gmin := n.Group.RmMin()
	gmax := n.Group.RmMax()
	fmt.Println(n.Group.Show())

	if n.left == nil {
		n.left = newNode(gmin)
	} else {
		n.left.insert(gmin)
	}
	if n.right == nil {
		n.right = newNode(gmax)
	} else {
		n.right.insert(gmax)
	}
	return nil
}

func (n *node) insert(k int) error {	
	n.Group.Validate()

	err := n.Group.Insert(k)	
	if n.Group.Full() {
		// if the group is full then do the magic split
		n.split()
	}
	return err
}

func (n *node) max() int {	
	if n.right != nil {
		return n.right.max()
	}
	return n.Group.Max()
}

// -------------------------------------------------------
type Jtree struct {
	root *node
}

func (j *Jtree) FixLeft() (err error) {
	if j.root.left == nil {
		return nil
	}

	// check to make sure the min of the group 
	// is larger than the max val in left branch	
	if j.root.left.Max() > j.root.Group.Min() {

		// if it not, remove the max val from the left branch
		max := j.root.left.RmMax()
		// remove the min val of the group
		min := j.root.Group.RmMin()
		// insert the left-branch-max-val into the group
		err = j.root.Group.Insert(max)
		// insert the group-min-val into the left branch
		err = j.root.left.insert(min)
	}		
	return err
}

func (j *Jtree) FixRight() (err error) {	
	if j.root.right == nil {
		return nil
	}

	// check to make sure the max of the group
	// is less than the min val of the right branch
	if j.root.Group.Max() > j.root.right.Min() {
		// if not, remove the min val from the right branch
		min := j.root.right.RmMin()
		// remove the max val of the group
		max := j.root.Group.RmMax()
		// insert the right-branch-min-val into the group
		err = j.root.Group.Insert(min)
		// insert tht group-max-val into the right branch
		err = j.root.right.insert(max)
	}		
	return 
}

func (j *Jtree) FixAll() (err error) { // redundant
	err = j.FixLeft()
	if err != nil {
		return err
	}
	err = j.FixRight()
	if err != nil {
		return err
	}
	return nil
}

func (j *Jtree) Insert(k int) (err error) {
	if j.root == nil {
		j.root = newNode(k)
		return nil
	} 

	err = j.FixAll()
	if err != nil {
		return err
	}

	err = j.root.insert(k)

	// err = j.FixAll()
	// if err != nil {
	// 	return err
	// }

	return err
}

func main() {
	jt := Jtree{}
	for i:=0; i<4; i++ {
		err := jt.Insert(i)
		
		if err != nil {
			fmt.Println(err)
		}
			//fmt.Println(jt.Show())
	}
	
	fmt.Println(jt.Show())
}



