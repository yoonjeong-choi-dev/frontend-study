#!/bin/bash

echo "Download a file from tftp server"
tftp <<EOF
mode binary
connect 127.0.0.1 7166
get server-resource.svg
quit
EOF

echo "Downloaded server-resource.svg"

echo "Validate Checksum"
go build .

./checksum-cli server-resource.svg
./checksum-cli "../tftp-server/main/payload.svg"

rm checksum-cli
rm server-resource.svg