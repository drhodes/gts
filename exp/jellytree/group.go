package main

import (
	"log"
	"fmt"
	"errors"
)


type Group struct {
	arr [3]int
	size int
}

func NewGroup(val int) Group {
	g := Group{[3]int{}, 0}
	g.Insert(val)
	return g
}

func (g *Group) Validate() {
	if g.size < 0 || g.size > 3 {
		log.Panic("size is wrong")
	}
}

func (g *Group) Show() string {
	g.Validate()
	return fmt.Sprint(g.arr[0:g.size])
}

func (g *Group) Max() int {
	g.Validate()
	return g.arr[g.size-1]
}

func (g *Group) Min() int {
	g.Validate()
	return g.arr[0]
}

func (g *Group) Full() bool {
	g.Validate()
	return g.size == 3
}

func (g *Group) RmMin() int {
	g.Validate()
	tmp := g.arr[0]
	g.arr[0] = g.arr[1]
	g.arr[1] = g.arr[2]
	g.size--
	return tmp
}

func (g *Group) RmMax() int {
	g.Validate()
	tmp := 0
	switch g.size {
	case 0: log.Panic("Trying to RmMax from an empty group")
	case 1: 
		tmp = g.arr[0]
		g.arr[1] = -1
		g.arr[2] = -1
	case 2:
		tmp = g.arr[1]
		g.arr[0] = g.arr[1]
		g.arr[2] = -1
	case 3:
		tmp = g.arr[2]
		g.arr[0] = g.arr[1]
		g.arr[1] = g.arr[2]
	default:
		log.Panic("g.size while in Group.RmMax is greater than 3") 
	}
	g.size--
	return tmp
}

func (g *Group) Insert(n int) error {
	g.Validate()

	switch g.size {
	case 1:
		g.size++
		switch {
		case n < g.arr[0]:
			g.arr[1] = g.arr[0]
			g.arr[0] = n			
			return nil
		case n > g.arr[0]:
			g.arr[1] = n
			return nil
		case n == g.arr[0]:
			msg := "1: We found two number with the same value, no repeats allowed"
			return errors.New(msg)
		}
	case 2:
		g.size++
		switch {
		case n < g.arr[0]:
			g.arr[2] = g.arr[1]
			g.arr[1] = g.arr[0]
			g.arr[0] = n
			return nil
		case n > g.arr[0]:
			g.arr[2] = n
			return nil
		case n == g.arr[0]:
			msg := "2: We found two number with the same value, no repeats allowed"
			return errors.New(msg)
		}
		
	case 0:
		g.size++
		g.arr[0] = n
		return nil
	default:
		msg := "You're trying to insert %d into a full group"
		return errors.New(msg)
	}
	return errors.New("Reached dead code in Group.Insert")
}
