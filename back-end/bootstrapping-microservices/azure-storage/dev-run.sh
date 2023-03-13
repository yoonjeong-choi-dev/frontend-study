#!/bin/bash

# export azure config
export $(xargs < .env.development)
export PORT=3000

npm run start:dev
