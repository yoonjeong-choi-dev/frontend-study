#!/bin/bash

# install
#go get github.com/axw/gocov/gocov
#go install github.com/axw/gocov/gocov
#go get github.com/smartystreets/goconvey
#go install github.com/smartystreets/goconvey

gocov test | gocov report
goconvey