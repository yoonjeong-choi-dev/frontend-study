#!/bin/bash

go build -o client

echo "Default Setting"
./client
echo ""

echo "Valid Query"
./client localhost:50051 '{"email": "yoonjeong@choi", "id": "yj"}'
echo ""

echo "Invalid Query 1"
./client localhost:50051 '{"email": "yoonjeong", "id": "yj"}'
echo ""

echo "Invalid Query 2"
./client localhost:50051 '{}'
echo ""

rm client