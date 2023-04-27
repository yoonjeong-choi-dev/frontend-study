FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0
#    GOOS=linux \
#    GOARCH=amd64

WORKDIR /build
COPY example/ example/
COPY go.mod go.sum config.go exec-example.go ./
RUN go mod download
RUN go build -o main example/example.go

RUN chmod +x main

FROM alpine

COPY --from=builder /build/main /app/main

CMD ["/app/main"]