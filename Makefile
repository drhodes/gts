
all:
	go build
	go build ./stack

run: all
	./gts -pkg main -type Point -gen stack

test:
	go test ./stack

clean: 	
	rm -f *~	
	go clean
	rm -f ./stack/*~
	go clean ./stack		

commit: clean
	git add .
	git commit -a
