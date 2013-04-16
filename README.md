An experiment in providing generic data structures through a simple template scheme.

  stack   Stack α
          A slice backed stack, it should beat the linked stack where 
		  there is rapid oscillating size. 

  rbtree  RedBlackTree α β
		  A red black tree ported from a java implementation provided with
          a talk given by Dr. Robert Sedgewick.
		  www.cs.princeton.edu/~rs/talks/LLRB/08Dagstuhl/Java/RedBlackBST.java

		  there is another implementation @ https://github.com/petar/GoLLRB
		  it uses interface{}s and type assertions, which is arguably better
		  or arguably worse.  We shall see!



TODO: 
	  benchmarks
	  linked stack



