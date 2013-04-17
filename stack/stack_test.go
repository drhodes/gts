package stack

import (
	"testing"
	"log"
)

func TestPop(t *testing.T) {
	s := NewStackα()
	
	s.Push(α(1))
	v, err := s.Pop()
	if err != nil {
		log.Println(err, "Pop Fails")
		t.Fail()
	}
	
	if v != α(1) {
		log.Println("Pop didn't match", v)
		t.Fail()
	}
}

func TestPointerPop(t *testing.T) {
	s := NewStackα()

	for i:=0; i<=1000; i++ {
		s.Push(α(i))
	}

	for i:=1000; i>=0; i-- {
		vp := s.PointerPop()
		if vp == nil {
			log.Println("Popped empty stack")
			t.Fail()
		}
		if *vp != α(i) {
			log.Println("Pop didn't match", *vp, i)
			t.Fail()
		}
	}

	// should be empty again
	if !s.Empty() {
		log.Println("Stack reports wrong emptiness")
		t.Fail()
	}
}


func TestMany(t *testing.T) {
	s := NewStackα()

	// should be empty
	if !s.Empty() {
		log.Println("Stack reports wrong emptiness")
		t.Fail()
	}

	for i:=0; i<=1000; i++ {
		s.Push(α(i))
	}

	for i:=1000; i>=0; i-- {
		v, err := s.Pop()
		if err != nil {
			log.Println(err, "Pop Fails")
			t.Fail()
		}

		if v != α(i) {
			log.Println("Pop didn't match", v, i)
			t.Fail()
		}
	}

	// should be empty again
	if !s.Empty() {
		log.Println("Stack reports wrong emptiness")
		t.Fail()
	}
}




