#!/bin/bash

HOST="http://localhost:7166"

curl $HOST?next=second
echo ""

curl $HOST?next=third
echo ""

curl $HOST?next=second
echo ""

curl $HOST?next=first
echo ""