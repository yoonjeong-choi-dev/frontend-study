#!/bin/bash

# MUST RUN
# 1. go build -o main.exe main.go
# 2. func start

HOST="http://localhost:7071/api/CustomHandlerFunction"

echo "Example 1. Empty Payload"
curl -X POST -H "Content-Type: application/json" --data '{}' \
  $HOST
echo ""

echo "Example 2. Recent Payload"
curl -X POST -H "Content-Type: application/json" --data '{"previous":true}' \
  $HOST
echo ""

echo "Example 3. All Payload"
curl -X POST -H "Content-Type: application/json" --data '{"all":true}' \
  $HOST
echo ""