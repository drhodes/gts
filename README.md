NOT READY FOR USE

An experiment in providing generic data structures 
through a simple template scheme.

How much overhead does a type assertion add? Quite a bit, depending on the context. For architectural members, or structures external to tight loops the overhead is irrelevant.  

So I wrote some benchmarks to highlight the widest performance disparity and _cast the worst possible light of the type assertions answer to generics_.  One might call this _dastardly unproportional and misrepresentative_.  If you would like to make fun of me, the code for these tests can be found in ./exp/bench.  There is a Makefile; make test.

    // roughly
    func(x) {
	    x * x
	}
BenchmarkInt	           2000000000  0.71 ns/op

    // roughly
    func() {
        x := interface_val.(int)
        x * x
    }
BenchmarkInterfaceUnsafe   100000000   12.5 ns/op
BenchmarkInterfaceSafe	   100000000   12.1 ns/op
BenchmarkTypeSwitch		   100000000   12.1 ns/op


## OFFERINGS:

------------------------------------------------------------------------------
# stack: Stack α

 A slice backed stack, it should beat the linked stack when the size oscillates rapidly. it's immune to thrashing.

-----------------------------------------------------------------------------
# rbtree: RedBlackTree α β 

A red black tree ported from a java implementation provided with a talk given by Dr. Robert Sedgewick:

www.cs.princeton.edu/~rs/talks/LLRB/08Dagstuhl/Java/RedBlackBST.java

Note: There is another golang implementation that has had some industrial use.  https://github.com/petar/GoLLRB


# TODO:
* benchmarks
* add safety token at the top of gts files to save data.

# CONSIDER:
* set
* linked stack
* skip list
* trie
* trie3
* treap

# Mangler:
* Package name should be unique.

# Merger:
* combine all rendered gts source into one file.
- 
  
  
