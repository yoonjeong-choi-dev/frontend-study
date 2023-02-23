#!/bin/bash

# Target: https://github.com/Tanmay-Teaches/golang/tree/master/chapter6/example3
# build script with M1 Process
docker build --platform linux/amd64 --tag go-for-backend-tetris-ai:1.0 .