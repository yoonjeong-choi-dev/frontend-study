#!/bin/bash

set -o allexport
source .env
set +o allexport


terraform destroy \
  -var="client_id=$CLIENT_ID" \
  -var="client_secret=${CLIENT_SECRET}"
