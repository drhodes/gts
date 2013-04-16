package main

import "testing"

func TestGroup(t *testing.T) {
	g := NewGroup(1)
	g.Insert(2)
	g.Insert(3)
	
	if g.arr[0] != 1 { t.Fail() }
	if g.arr[1] != 2 { t.Fail() }
	if g.arr[2] != 3 { t.Fail() }
	
	if g.size != 3 { t.Fail() }

	v := g.RmMin() 
	if v != 1 {t.Fail()}
	if g.size != 2 { t.Fail() }

	v = g.RmMin() 
	if v != 2 {t.Fail()}
	if g.size != 1 { t.Fail() }

	g.Insert(6)
	if g.RmMax() != 6 { t.Fail() }
	if g.size != 1 { t.Fail() }

	g.Insert(-1)
	if g.RmMin() != -1 { t.Fail() }
	if g.size != 1 { t.Fatal(g) }
	
	


}