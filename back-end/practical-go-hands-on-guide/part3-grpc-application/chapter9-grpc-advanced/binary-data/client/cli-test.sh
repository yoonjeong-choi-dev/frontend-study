#!/bin/bash

# Exercise 9.2
go build .

echo "Simple Test"
./grpc-cli --file-path ./go.mod localhost:50051
echo ""

echo "With Custom Name"
./grpc-cli --file-path ./go.mod --user-name yj localhost:50051
echo ""

echo "With Full Option"
./grpc-cli --file-path ./go.mod --user-name yjchoi --file-name test.txt localhost:50051
echo ""


rm ./grpc-cli