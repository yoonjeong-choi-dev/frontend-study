#!/bin/bash

HOST="http://localhost:7166"

echo "Example 1. Get with no query"
curl $HOST/hello
echo ""

echo "Example 2. Get with name query"
curl $HOST/hello?name=YJCHOI
echo ""

echo "Example 3. Post with empty data"
curl $HOST/greet -X POST
echo ""

echo "Example 4. Post with full data"
curl $HOST/greet -X POST \
  -d "name=yoonejeong&greeting=Hello"
echo ""

echo "Example 5. Post with wrong data"
curl $HOST/greet -X POST -d "name=yj;"
echo ""
