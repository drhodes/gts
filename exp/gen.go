package main

// 1 proper superset.

type List(a Eq) struct {
	val a
	next *a
}

func (l *List(a)) Map(f func(a)a) {
	l.val = f(l.val)
	if next != nil {
		l.next.Map(f)
	}
}

func (l *List(a)) Eq(other *List(a)) {
	for {
		if l.val != other.val {
			return false
		}		
	}
	return true
}
