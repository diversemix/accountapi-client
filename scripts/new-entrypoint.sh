#!/usr/bin/dumb-init /bin/sh
set -e

echo "*************************** PATCHED"

function pgping() {

  PROTOCOL_VERSION="\x00\x03\x00\x00"
  COMMAND="user"

  NC_TIMEOUT=3

  POSTGRES_SERVER=postgresql
  POSTGRES_PORT=5432
  USERNAME="interview_accountapi_user"

  PACKET_SIZE="\x00\x00\x00\x$(printf '%02x' $((
    4 +
    ${#PROTOCOL_VERSION} / 4 +
    ${#COMMAND} +
    1 +
    ${#USERNAME} +
    2
  )))"

  test "$(
    echo -ne "${PACKET_SIZE}${PROTOCOL_VERSION}${COMMAND}\x00${USERNAME}\x00\x00" |
    nc -w $NC_TIMEOUT $POSTGRES_SERVER $POSTGRES_PORT 2>/dev/null | head -c1
  )" == "R"

  if [ $? -eq 0 ]; then
    echo "health check passed"
    return 0
  else
    echo "health check failed"
    return 1
  fi
}

until pgping
do 
    echo "retrying..."
    sleep 1
done

# This is fragile as it depends on us knowing the entrypoint
/app/entrypoint.sh
