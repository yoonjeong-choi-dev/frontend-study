#!/bin/bash

HOST="http://localhost:7166"

echo "Example 1. Empty Payload"
curl $HOST -X POST \
  -d "{}"
echo ""

echo "Example 2. Valid Payload"
curl $HOST -X POST \
  -d '{"name":"yj-choi", "age": 31}'
echo ""

echo "Example 3. Invalid Payload (1)"
curl $HOST -X POST \
  -d '{"name":"yj-choi", "age": -123}'
echo ""

echo "Example 4. Invalid Payload (2)"
curl $HOST -X POST \
  -d '{"age": 123}'
echo ""