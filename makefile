dev:
	gin run main.go 

build:
	go build main.go

run:
	go run main.go

test:
	go test -v -coverprofile cover.out ./package/...

test-cover:
	go tool cover -html=cover.out -o cover.html