#!/bin/bash

set -o allexport
source .env
set +o allexport

terraform apply \
  -var="client_id=$CLIENT_ID" \
  -var="client_secret=${CLIENT_SECRET}" \
  -var="storage_account_name=${AZURE_STORAGE_ACCOUNT_NAME}" \
  -var="storage_access_key=${AZURE_STORAGE_ACCESS_KEY}" \
  -auto-approve
