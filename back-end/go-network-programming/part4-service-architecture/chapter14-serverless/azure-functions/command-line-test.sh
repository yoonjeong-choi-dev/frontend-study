#!/bin/bash

# Deploy 프로세스마다 해당 URL은 달라짐
HOST="https://yj-azure-serverless.azurewebsites.net/api/customhandlerfunction"

echo "Example 1. Empty Payload"
curl -X POST -H "Content-Type: application/json" --data '{}' \
  $HOST
echo ""

echo "Example 2. Recent Payload"
curl -X POST -H "Content-Type: application/json" --data '{"previous":true}' \
  $HOST -v
echo ""

echo "Example 3. All Payload"
curl -X POST -H "Content-Type: application/json" --data '{"all":true}' \
  $HOST -v
echo ""