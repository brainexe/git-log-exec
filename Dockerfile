FROM golang:1.15-alpine as builder

WORKDIR /code/
COPY . ./

RUN go build -ldflags="-s -w" -o /git-log-exec *.go

FROM alpine:latest as alpine
RUN apk add --no-cache git bash ca-certificates
COPY --from=builder git-log-exec /

COPY docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["--help"]
