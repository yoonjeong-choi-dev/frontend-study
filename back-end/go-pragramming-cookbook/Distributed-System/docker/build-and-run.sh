#!/bin/bash

pushd server
# {package name}.{variable} 형태로 빌드 시 정보 전달 가능
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.version=1.0 -X main.buildDate=$(date +%s)"
popd
docker build -t go-simple-server .
docker run -d -p 7166:4000 go-simple-server
