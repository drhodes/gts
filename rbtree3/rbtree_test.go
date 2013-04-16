package rbtree

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	rb := RedBlackBST{}

	for i := α(-100); i < 100; i++ {
		rb.put(i, β(i*2))
	}

	for i := α(-100); i < 100; i++ {
		v, ok := rb.get(i)
		if ok {
			if v != β(i*2) {
				t.Fail()
			}
		} else {
			t.Fail()
		}
	}

	for i := -100; i < 100; i++ {
		rb.Delete(α(i))
	}
}

func TestDelete(t *testing.T) {
	rb := RedBlackBST{}

	for i := α(-100); i < 100; i++ {
		rb.put(i, β(i*2))
	}
	
	for i := -100; i < 100; i++ {
		rb.Delete(α(i))
	}

	if rb.size() != 0 {
		log.Println("Delete test not working")
		t.Fail()
	}
}
