package stack

import (
	"errors"
)

// dummy start --------------------------------------------
func Zero𝞃() 𝞃 {
	return 𝞃(0)
}
// dummy end ----------------------------------------------

type 𝞃 int

// array based stack with repeated doubling
type Stack𝞃 struct {	
	arr []𝞃
	cur int
}

func NewStack𝞃() Stack𝞃 {
	return Stack𝞃{make([]𝞃, 1), 0}
	
}

func (s *Stack𝞃) Push(el 𝞃) {
	s.maybeGrow()
	s.arr[s.cur] = el
	s.cur++
}

func (s *Stack𝞃) Empty() bool {
	return s.cur == 0 
}

// alternatively this could return a pointer or panic (but that seems excessive).
func (s *Stack𝞃) Pop() (𝞃, error) {
	if s.Empty() {
		return Zero𝞃(), errors.New("Can't pop stack, it's already empty")
	}
	el := s.arr[s.cur - 1]		
	s.cur--
	s.maybeShrink()
	return el, nil
}

// less allocation, but the user needs to check for nil.
func (s *Stack𝞃) PointerPop() *𝞃 {
	if s.Empty() {
		return nil
	}
	el := s.arr[s.cur - 1]
	s.cur--
	s.maybeShrink()
	return &el
}

// Half the stack capacity if cur index is quarter the length
func (s *Stack𝞃) maybeShrink() {	
	if s.cur <= len(s.arr) / 4 {
		arr := make([]𝞃, len(s.arr) / 2)
		copy(arr, s.arr[0:s.cur+1])
		s.arr = arr
	}
}

// Double the stack capacity if we're out of space.
func (s *Stack𝞃) maybeGrow() {
	if s.cur == len(s.arr) - 1 {
		// log.Println("Growing", s.cur, len(s.arr))
		// create a new slice, twice the size
		arr := make([]𝞃, len(s.arr) * 2)
		copy(arr, s.arr)		
		s.arr = arr
	}
}
