package assert

import (
	"testing"
	"log"
)


func squareInt(x int) int {
	return x * x
}

func squareInterfaceUnsafe(x interface{}) int {
	tmp := x.(int)
	return tmp * tmp
}

func squareInterfaceSafe(x interface{}) int {
	tmp, ok := x.(int)
	if ok {
		return tmp*tmp
	} 
	log.Panic("aw crap")
	return -1
}

func squareTypeSwitch(x interface{}) int {
	switch x.(type) {
	case int: 
		tmp := x.(int)
		return tmp * tmp
	}
	log.Panic("aw crap")
	return -1
}



func BenchmarkInt(b *testing.B) {
	for i:=0; i<b.N; i++ {
		squareInt(3)
	}
}


func BenchmarkInterfaceUnsafe(b *testing.B) {
	v := interface{}(3)
	for i:=0; i<b.N; i++ {
		squareInterfaceUnsafe(v)
	}
}

func BenchmarkInterfaceSafe(b *testing.B) {
	//b.StartTimer()
	v := interface{}(3)
	for i:=0; i<b.N; i++ {
		squareInterfaceSafe(v)
	}
	//b.StopTimer()
}

func BenchmarkTypeSwitch(b *testing.B) {
	v := interface{}(3)
	for i:=0; i<b.N; i++ {
		squareInterfaceSafe(v)
	}
}
