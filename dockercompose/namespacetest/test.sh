#!/bin/bash
set -e

echo "injected endpoint is $ENDPOINT"

STATUS=`curl -X GET -w '%{http_code}' -o response.txt -s $ENDPOINT`

echo "curl $ENDPOINT returned status $STATUS"
if [[ $STATUS -ne 200 ]]; then
    exit 1
fi