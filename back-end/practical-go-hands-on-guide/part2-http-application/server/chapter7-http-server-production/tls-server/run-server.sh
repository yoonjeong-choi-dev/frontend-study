#!/bin/bash

go build -o server

cert_fil_path="./certification"
export TLS_CERT_FILE_PATH=${cert_fil_path}/server.crt
export TLS_KEY_FILE_PATH=${cert_fil_path}/server.key

./server

