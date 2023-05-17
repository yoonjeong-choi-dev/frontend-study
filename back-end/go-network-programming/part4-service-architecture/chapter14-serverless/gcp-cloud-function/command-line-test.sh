#!/bin/bash

# Deploy 프로세스마다 해당 URL은 달라짐
HOST="https://us-central1-yj-go-network-serverless.cloudfunctions.net/MainHandler"

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