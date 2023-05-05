#!/bin/bash

GOMAXPROCS=1 go test -bench . -benchmem -benchtime=1s