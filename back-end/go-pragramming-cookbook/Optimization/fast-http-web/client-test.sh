#!/bin/bash

ITEM_HOST="http://localhost:7166/item"

curl $ITEM_HOST/test1 -X POST
curl $ITEM_HOST/test2 -X POST
curl $ITEM_HOST/test3 -X POST

curl $ITEM_HOST
echo ""