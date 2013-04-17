An experiment in providing generic data structures through a simple template scheme.

  stack   Stack α
          A slice backed stack, it should beat the linked stack when
		  the stack size oscillates rapidly. it's immune to thrashing.

  rbtree  RedBlackTree α β
		  A red black tree ported from a java implementation provided with
          a talk given by Dr. Robert Sedgewick.
		  www.cs.princeton.edu/~rs/talks/LLRB/08Dagstuhl/Java/RedBlackBST.java

		  Note:
		  there is another implementation @ https://github.com/petar/GoLLRB
		  It uses interface{}s and type assertions, which is arguably better
		  or arguably worse. 


TODO: 
	  benchmarks

CONSIDER:
	  linked stack
	  skip list
	  skip tree
	  trie
	  trie3


