FROM golang:latest AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /home/ubuntu/projects/super-simple-api-go-orm
COPY go.mod .
RUN go mod download
COPY . .
RUN go build main.go

FROM scratch
COPY --from=builder /go/src .
ENTRYPOINT  ["./main"]