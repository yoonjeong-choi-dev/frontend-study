#!/bin/bash

go build main.go

# http -> https 로 리다이렉션 해줌
./main http://github.com

# remove execution file
rm ./main