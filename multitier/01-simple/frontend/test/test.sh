#!/bin/bash
set -e

echo "injected endpoint is $ENDPOINT"

STATUS=`curl -X GET -w '%{http_code}' -o response.txt -s $ENDPOINT`

RESPONSE=`cat response.txt`

if [[ $STATUS -ne 200 ]]; then
    exit 1
fi

if [[ "$RESPONSE" != *"Namespace: example of a Vite+React frontend"* ]]; then
    echo "Unexpected response: $RESPONSE"
    exit 1
fi
