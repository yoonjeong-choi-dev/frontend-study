#!/bin/bash

curl --request "POST" \
  --location "http://localhost:7166/twirp/GreeterService/Greet" \
  --header "Content-Type:application/json" \
  --data '{"greeting":"This is from curl", "name":"yj-curl-cli"}'

echo ""