#!/bin/bash

go build .

trap "rm main" INT

./main