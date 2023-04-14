#!/bin/bash

go build .
trap "rm main" SIGTERM SIGINT

echo "Example 1. golang.org"
./main golang.org

echo "Example 2. google.com"
./main google.com

rm main