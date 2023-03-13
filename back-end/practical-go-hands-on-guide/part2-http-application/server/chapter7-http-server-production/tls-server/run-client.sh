#!/bin/bash

echo "Example 1. curl with http"
curl localhost:7166

echo ""
echo "Example 2. cur with https without cert file"
curl https://localhost:7166

echo ""
echo "Example 3. cur with https with cert file"
curl --cacert ./certification/server.crt https://localhost:7166/api

echo ""