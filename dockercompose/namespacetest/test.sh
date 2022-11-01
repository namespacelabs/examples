#!/bin/bash
set -e

# "-r" removes quotes from the output.
ENDPOINT=`cat /namespace/config/runtime.json | jq -r ".stack_entry[0].service[0].endpoint"`

echo "Extracted endpoint $ENDPOINT"

STATUS=`curl -X POST -w '%{http_code}' -o response.txt --data '{"name": "'testItem'"}' $ENDPOINT/add`
if [[ $STATUS -ne 200 ]]; then
    echo "failed to add item (status $STATUS)"
    exit 1
fi

echo "SUCCESS!"