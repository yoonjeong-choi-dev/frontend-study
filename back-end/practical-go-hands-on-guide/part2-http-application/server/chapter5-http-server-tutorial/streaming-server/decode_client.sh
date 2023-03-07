#!/bin/bash

echo 'Valid Log Request'
curl -X POST http://localhost:7166/decode \
  -d '
  {"user_ip": "172.121.19.21", "event": "click_on_add_cart"}
  {"user_ip": "172.121.19.22", "event": "click_on_checkout"}
  '

echo ''
echo 'Invalid Log Request'
curl -X POST http://localhost:7166/decode \
  -d '
  {"user_ip": "172.121.19.21", "event": "click_on_add_cart"}
  {"user_ip": "172.121.19.22", "event": 123}
  {"user_ip": "172.121.19.22", "event": "click_on_checkout", "user_data":"additional field"}
  '
echo ''
