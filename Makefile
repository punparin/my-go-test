test:
	CGO_ENABLED=0 go test -v -covermode=count .

run:
	CGO_ENABLED=0  go run .
