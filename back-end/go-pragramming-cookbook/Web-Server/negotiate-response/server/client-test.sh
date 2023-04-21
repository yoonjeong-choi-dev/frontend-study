#!/bin/bash

HOST=http://localhost:7166

echo "Example 1. text/xml type"
curl $HOST -H "Content-Type: text/xml"
echo ""

echo "Example 2. application/json type"
curl $HOST -H "Content-Type: application/json"
echo ""
