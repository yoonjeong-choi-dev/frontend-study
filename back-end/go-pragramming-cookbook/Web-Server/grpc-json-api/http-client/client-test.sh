#!/bin/bash

HOST="http://localhost:7166"

echo "Example 1. Set (key1, val1)"
curl $HOST/set -X POST \
  -d '{"key":"key1", "value":"va1"}'
echo ""

echo "Example 2. Get key1"
curl $HOST/get?key=key1
echo ""

echo "Example 3. Get badkey"
curl $HOST/get?key=badkey -v
echo ""
