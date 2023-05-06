build:
	# use @ to hide output if needed 
	go build -o bin/ecolabel

run: build
	./bin/ecolabel

test:
	go test -v ./...

bench:
	go test -bench=. -benchmem ./internal/...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

cover-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	xdg-open coverage.html || open coverage.html

