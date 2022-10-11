#!/bin/bash

# "-r" removes quotes from the output.
ENDPOINT=`cat /namespace/config/runtime.json | jq -r ".stack_entry[0].service[0].endpoint"`

for NAME in item1 item2
do
    RESPONSE=`curl -s -X POST --data '{"name": "'$NAME'"}' $ENDPOINT/add`
    if [[ ! -z "$RESPONSE" ]]; then
        echo "failed to add $NAME: $RESPONSE"
        exit 1
    fi
done

RESPONSE=`curl -s $ENDPOINT/list`
echo "Got list response:\n$RESPONSE"

# "-c" compacts the results to a single line.
LIST=`echo $RESPONSE | jq -c 'map(.name)'`
EXPECTED="[\"item1\",\"item2\"]"
if [[ "$LIST" != "$EXPECTED" ]]; then
    echo "Unexpected result: $LIST"
    exit 1
fi

