#!/bin/bash

HOST="http://localhost:7166"

echo "Example 1. Set Value"
curl $HOST/set -X POST \
  -d "value=test-value-to-save"
echo ""

echo "Example 2. Get with default"
curl $HOST/get/default
echo ""

echo "Example 3. Get data from storage"
curl $HOST/get
echo ""

