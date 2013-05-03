


mangle:
	./gts -mangle ./rbtree

build: 
	go build
	go build ./stack
	go build ./rbtree
	$(MAKE) mangle

run: all
	./gts -pkg main -type Point -gen stack

install:
	go install

test:
	@go test ./stack
	@go test ./rbtree

clean: 	
	rm -f *~	
	go clean
	rm -f ./stack/*~
	go clean ./stack		
	rm -f ./rbtree/*~
	go clean ./rbtree

commit: clean
	git add .
	git commit -a
