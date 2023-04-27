#!/bin/bash

echo "Get version:"
curl http://localhost:7166/version
echo ""

echo "Clean up..."
docker stop `docker ps -a -q`
docker rm `docker ps -a -q`

pushd server
rm server
popd