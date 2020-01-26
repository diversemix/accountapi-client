#!/bin/bash

echo Waiting for ${URL_UNDER_TEST}

until [ $(curl -f -s "${URL_UNDER_TEST}/v1/health") ]
do  
    sleep 1
    echo "retrying..."
done
