

all:
	go build

perf: all
	rm -f perf.data
	perf record ./jellytree
	perf report
	rm -f perf.data
