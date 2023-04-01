#!/bin/bash

go build .

echo "Example input 1: test case"
echo "test case" | ./unix-pipe

echo "Example input 1: test case case Test"
echo "test case case Test" | ./unix-pipe

rm unix-pipe