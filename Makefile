run-tests:
	go test -race -v ./...

precommit:
	go fmt ./...