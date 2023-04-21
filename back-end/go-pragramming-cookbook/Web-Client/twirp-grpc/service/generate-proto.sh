#!/bin/bash

# install twirp plugin
#go install github.com/twitchtv/twirp/protoc-gen-twirp@latest

protoc --go_out=. --go_opt=paths=source_relative --twirp_out=. --twirp_opt=paths=source_relative greeter.proto
