#!/bin/bash

mkdir internal

mockgen -destination internal/mocks.go -package internal -source=interface.go GetSetter