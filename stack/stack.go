package stack

import (
	"errors"
	"log"
)

// gts instructions
// Implement: func Zeroα() α
// It's a default constructor used to appease the typesystem
// in the event of multiple return with error e.g:
//   return Zeroα(), err

// dummy start --------------------------------------------
func Zeroα() α {
	return α(0)
}

type α int
// dummy end ----------------------------------------------


// array based stack with repeated doubling
type Stackα struct {	
	arr []α
	cur int
}

func NewStackα() Stackα {
	return Stackα{make([]α, 1), 0}
	
}

func (s *Stackα) Push(el α) {
	s.maybeGrow()
	s.arr[s.cur] = el
	s.cur++
}

func (s *Stackα) Empty() bool {
	return s.cur == 0 
}

// alternatively this could return a pointer or panic (but that seems excessive).
func (s *Stackα) Pop() (α, error) {
	if s.Empty() {
		return Zeroα(), errors.New("Can't pop stack, it's already empty")
	}
	el := s.arr[s.cur - 1]		
	s.cur--
	s.maybeShrink()
	return el, nil
}

// less allocation, but the user needs to check for nil.
func (s *Stackα) PointerPop() *α {
	if s.Empty() {
		return nil
	}
	el := s.arr[s.cur - 1]
	s.cur--
	s.maybeShrink()
	return &el
}

// less allocation, but the user needs to check for nil.
func (s *Stackα) PanicPop() α {
	if s.Empty() {
		log.Panic("Trying to pop an empty stack")
	}
	el := s.arr[s.cur - 1]
	s.cur--
	s.maybeShrink()
	return el
}

// Half the stack capacity if cur index is quarter the length
func (s *Stackα) maybeShrink() {	
	if s.cur <= len(s.arr) / 4 {
		arr := make([]α, len(s.arr) / 2)
		copy(arr, s.arr[0:s.cur+1])
		s.arr = arr
	}
}

// Double the stack capacity if we're out of space.
func (s *Stackα) maybeGrow() {
	if s.cur == len(s.arr) - 1 {
		// log.Println("Growing", s.cur, len(s.arr))
		// create a new slice, twice the size
		arr := make([]α, len(s.arr) * 2)
		copy(arr, s.arr)		
		s.arr = arr
	}
}
