#!/bin/bash

echo 'Download simple text file'
curl localhost:7166/download?filename=data1.txt
echo ''

echo 'Download json file'
curl localhost:7166/download?filename=data2.json
echo ''