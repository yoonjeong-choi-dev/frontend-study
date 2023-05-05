#!/bin/bash

for var in {1..5}
do
  curl http://localhost:7166/guess?message=test
  echo ""
  curl http://localhost:7166/guess?message=password
  echo ""
done

