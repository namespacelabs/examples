#!/bin/bash

# "-r" removes quotes from the output.
ENDPOINT=`cat /namespace/config/runtime.json | jq -r ".stack_entry[0].service[0].endpoint"`

RESPONSE=`curl -s -X POST --data '{"name": "item1"}' $ENDPOINT/add`
if [[ ! -z "$RESPONSE" ]]; then
    echo "failed to add item1: $RESPONSE"
    exit 1
fi

RESPONSE=`curl -s -X POST --data '{"name": "item2"}' $ENDPOINT/add`
if [[ ! -z "$RESPONSE" ]]; then
    echo "failed to add item2: $RESPONSE"
    exit 1
fi

RESPONSE=`curl -s $ENDPOINT/list`

LEN=`echo $RESPONSE | jq length`
if [[ "$LEN" != "2" ]]; then
    echo "Expected two list items, got: $LEN\n$RESPONSE"
    exit 1
fi

NAME=`echo $RESPONSE | jq -r .[0].name`
if [[ "$NAME" != "item1" ]]; then
    echo "Unexpexted item: $NAME"
    exit 1
fi

NAME=`echo $RESPONSE | jq -r .[1].name`
if [[ "$NAME" != "item2" ]]; then
    echo "Unexpexted item: $NAME"
    exit 1
fi
