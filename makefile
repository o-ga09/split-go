run: 
	go run ./cmd/main.go -n $(arg)


test:
	go test ./...