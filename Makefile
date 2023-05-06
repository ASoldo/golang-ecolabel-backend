build:
	# use @ to hide output if needed 
	go build -o bin/ecolabel

run: build
	./bin/ecolabel

test:
	go test -v ./...

bench:
	go test -bench=. -benchmem ./internal/...
