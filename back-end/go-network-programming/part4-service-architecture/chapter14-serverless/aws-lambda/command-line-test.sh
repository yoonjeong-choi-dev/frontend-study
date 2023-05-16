#!/bin/bash

TEMP="response.json"

aws lambda invoke --function-name "yj-go-network-serverless-app" $TEMP
cat $TEMP
echo ""

aws lambda invoke --function-name "yj-go-network-serverless-app" --cli-binary-format "raw-in-base64-out" \
  --payload '{"previous":true}' $TEMP
cat $TEMP
echo ""

aws lambda invoke --function-name "yj-go-network-serverless-app" --cli-binary-format "raw-in-base64-out" \
  --payload '{"all":true}' $TEMP
cat $TEMP
echo ""

rm $TEMP