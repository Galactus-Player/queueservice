#!/bin/bash

# export GO_POST_PROCESS_FILE=goimports
export OPENAPI_YML=./api/openapi.yaml

mkdir -p client
openapi-generator generate --enable-post-process-file -o client/ -i $OPENAPI_YML -g typescript-fetch
