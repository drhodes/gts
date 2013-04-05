package stack

import (
	"testing"
	"log"
)


func TestPop(t *testing.T) {
	s := NewStack𝞃()
	
	s.Push(𝞃(1))
	v, err := s.Pop()
	if err != nil {
		log.Println(err, "Pop Fails")
		t.Fail()
	}
	
	if v != 𝞃(1) {
		log.Println("Pop didn't match", v)
		t.Fail()
	}
}

func TestPointerPop(t *testing.T) {
	s := NewStack𝞃()

	for i:=0; i<=1000; i++ {
		s.Push(𝞃(i))
	}

	for i:=1000; i>=0; i-- {
		vp := s.PointerPop()
		if vp == nil {
			log.Println("Popped empty stack")
			t.Fail()
		}
		if *vp != 𝞃(i) {
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
	s := NewStack𝞃()

	// should be empty
	if !s.Empty() {
		log.Println("Stack reports wrong emptiness")
		t.Fail()
	}

	for i:=0; i<=1000; i++ {
		s.Push(𝞃(i))
	}

	for i:=1000; i>=0; i-- {
		v, err := s.Pop()
		if err != nil {
			log.Println(err, "Pop Fails")
			t.Fail()
		}

		if v != 𝞃(i) {
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




