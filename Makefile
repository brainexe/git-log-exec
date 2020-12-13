
.PHONY: test docker-build

git-log-exec:
	go build -o git-log-exec *.go

test:
	go test ./... -race

docker-build:
	docker build . -t brainexe/git-log-exec