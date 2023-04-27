#!/bin/bash

HOST="http://localhost:7166"

curl $HOST/counter
echo ""

curl $HOST/timer
echo ""

curl $HOST/counter
echo ""

curl $HOST/report
echo ""