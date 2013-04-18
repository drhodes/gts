package gts


/*

-cu 
--custom-url

-cf
--custom-file 

--stdin  

--list 

-tp 
--type-params = Type1 Type2 ... TypeN

type-params checks the gts file in the repository 
to ensure that the correct number of type parameters were
provided.  If not, gts emits an error message and exits
with failure code

gts -tag 1 -g rbtree -tp Int 

*/
