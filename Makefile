
.PHONY: test docker-build

git-log-exec:
	go build -o git-log-exec *.go

test:
	go test ./...

docker-build:
	docker build . -t brainexe/git-log-exec