#!/bin/bash

set -o allexport
source .env
set +o allexport


terraform apply \
  -var="client_id=$CLIENT_ID" \
  -var="client_secret=${CLIENT_SECRET}" \
  -auto-approve
