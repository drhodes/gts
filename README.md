An experiment in providing generic data structures 
through a simple template scheme.

How much overhead does a type assertion add? Quite a bit, depending on the context. For architectural members, or structures external to tight loops the overhead is irrelevant.  

So I wrote some benchmarks and tried to highlight the widest disparity to cast the worst possible light of the type assertions answer to generics.  One might call this dastardly unproportional and misrepresentative.  The code for these tests can be found in ./exp/bench.  There is a Makefile; make test.

~~ func() { x * x }
BenchmarkInt	           2000000000  0.71 ns/op

vs

~~ func() {
~~    x := interface_val.(int)
~~    x * x
~~ }

BenchmarkInterfaceUnsafe   100000000   12.5 ns/op
BenchmarkInterfaceSafe	   100000000   12.1 ns/op
BenchmarkTypeSwitch		   100000000   12.1 ns/op


OFFERINGS:

-----------------------------------------------------------------------------
stack: Stack α

 A slice backed stack, it should beat the linked stack when the size oscillates rapidly. it's immune to thrashing.

-----------------------------------------------------------------------------
rbtree: RedBlackTree α β 

A red black tree ported from a java implementation provided with a talk given by Dr. Robert Sedgewick:

www.cs.princeton.edu/~rs/talks/LLRB/08Dagstuhl/Java/RedBlackBST.java

Note: There is another golang implementation that has had some industrial use.  https://github.com/petar/GoLLRB






TODO:
	  benchmarks

CONSIDER:
      set
	  linked stack
	  skip list
	  trie
	  trie3
	  generifying the stdlib containers

is the name mangler sufficient?


maybe it would be nice to generate one file instead of filling user dir. yes.
  easy: parse/combine imports
  ..
  
  