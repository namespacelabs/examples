#!/bin/bash
set -e

# "-r" removes quotes from the output.
ENDPOINT=`cat /namespace/config/runtime.json | jq -r ".stack_entry[0].service[0].endpoint"`

STATUS=`curl -X GET -w '%{http_code}' -o response.txt -s $ENDPOINT`

echo "curl $ENDPOINT returned status $STATUS"
if [[ $STATUS -ne 200 ]]; then
    exit 1
fi