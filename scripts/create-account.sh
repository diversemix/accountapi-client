#!/bin/bash

EP=localhost:8080/v1/organisation/accounts
ID=$(uuidgen)

echo ---- POST
DATA="{ \"data\": { \"id\": \"${ID}\", \"organisation_id\": \"d34d80b8-7b4b-4585-ba29-8460f21ca0db\", \"type\": \"accounts\", \"attributes\": {\"country\": \"GB\", \"bic\": \"NWBKGB22\", \"bank_id\":\"400302\"} } }"
curl -v -d "${DATA}" \
    -H "Accept: application/vnd.api+json" \
    -H "Content-Type: application/vnd.api+json" \
    ${EP}

echo ---- GET
curl ${EP}
