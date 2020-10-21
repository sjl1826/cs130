default: build

run: build
	go run ./cmd/cs130_back/main.go

build:
	go build -o cs130_back.out ./cmd/cs130_back/main.go

clean:
	rm -rf cs130_back.out

test:
	go test -v ./...