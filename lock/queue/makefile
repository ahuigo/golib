bench:
	go test -v -bench=.
s:
	go build -o tmp queue.go
	go tool objdump -s .Queue.PushTailCAS tmp
	go tool objdump -s .Queue.PushTailMutex tmp
a:
	go build -o tmp a.go
	go tool objdump tmp
