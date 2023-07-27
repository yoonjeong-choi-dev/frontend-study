#!/bin/bash

go run generate-cert-cli.go -cert serverCert.pem -key serverKey.pem -host localhost
go run generate-cert-cli.go -cert clientCert.pem -key clientKey.pem -host localhost