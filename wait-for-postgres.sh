#!/bin/bash
#
# wait-for-postgres.sh hostname [port] cmd...
#
# After ensuring that we can connect to Postgres on the given hostname:port combination, run cmd...
# port is optional. If the 2nd parameter is a number, then it is treated as a port.
#
# WAIT_FOR_POSTGRES environment variable can be set to false to disable waiting.
# If WAIT_FOR_POSTGRES is set to false, the paramaters must STILL be in the correct format, the wait loop will just not be executed.

set -e

testparams() {
  if [[ "$#" -ne 0 ]]; then
    return 0
  else
    >&2 echo "Not enough parameters: wait-for-postgres.sh hostname [port] cmd..."
    return 1
  fi
}

testparams "$@"
if [[ "$?" -eq 1 ]] ; then
  exit 1
fi

host="$1"
shift

testparams "$@"
if [[ "$?" -eq 1 ]] ; then
    exit 2
fi

port=5432
isport='^[0-9][0-9]+$'
if [[ $1 =~ $isport ]]; then
  port=$1
  shift
fi

testparams "$@"
if [[ "$?" -eq 1 ]] ; then
    exit 2
fi

cmd="$@"

if [ -z "$WAIT_FOR_POSTGRES" ] || ([ -n "$WAIT_FOR_POSTGRES" ] &&
        ([ "$WAIT_FOR_POSTGRES" == "true" ] || [ "$WAIT_FOR_POSTGRES" == "True" ] || [ "$WAIT_FOR_POSTGRES" == "TRUE" ] || [ "$WAIT_FOR_POSTGRES" == "1" ])) ; then
  >&2 echo "WAIT_FOR_POSTGRES unset, or set to true, waiting..."
  until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$host" -p "$port" -U "postgres" -c '\q'; do
    >&2 echo "Postgres is unavailable on $host:$port password: $POSTGRES_PASSWORD - sleeping"
    sleep 1
  done
else
  >&2 echo "WAIT_FOR_POSTGRES set to false, proceeding."
fi

exec $cmd