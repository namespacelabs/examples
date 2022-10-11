#!/bin/bash

# "-r" removes quotes from the output.
ENDPOINT=`cat /namespace/config/runtime.json | jq -r ".stack_entry[0].service[0].endpoint"`

for NAME in item1 item2
do
    STATUS=`curl -s -X POST -w '%{http_code}' -o response.txt --data '{"name": "'$NAME'"}' $ENDPOINT/add`
    if [[ $STATUS -ne 200 ]]; then
        echo "failed to add $NAME (status $STATUS)"
        cat response.txt
        exit 1
    fi
done

STATUS=`curl -s -w '%{http_code}' -o response.txt $ENDPOINT/list`
if [[ $STATUS -ne 200 ]]; then
    echo "failed to list items (status $STATUS)"
    cat response.txt
    exit 1
fi

# "-c" compacts the results to a single line.
LIST=`cat response.txt | jq -c 'map(.name)'`
EXPECTED="[\"item1\",\"item2\"]"
if [[ "$LIST" != "$EXPECTED" ]]; then
    echo "Unexpected result: $LIST"
    exit 1
fi

