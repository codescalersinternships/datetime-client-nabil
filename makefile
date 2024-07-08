httpclient:
	go run ./cmd/

test:
	go test -v -cover ./...

format:
	go fmt ./...
