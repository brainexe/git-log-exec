
.PHONY: test

git-log-exec:
	go build -o git-log-exec *.go

test:
	go test ./...