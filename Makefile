
all:
	go build
	go build ./stack

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

commit: clean
	git add .
	git commit -a
