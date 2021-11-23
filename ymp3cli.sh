#!/bin/bash

# Variables.
URL="http://127.0.0.1:8000/"


# Prevents the server continuing to run if the terminal dies.
trap 'kill $(ps -o pid= --ppid $$) 2>/dev/null' exit

mkdir -p music


go run src/main.go 2 >&1 >/dev/null &
printf "Waiting for server to start"

# Wait for the server to start.
while true; do
  STATUS=$(curl -s -o /dev/null -w "%{http_code}\n" $URL)

    if [ $STATUS == 200 ] ; then
      echo ""
      echo "Server is up and running"

      sleep 2

      break
    fi

  printf '.'
  sleep 0.1
done


deno run --allow-env --allow-net --allow-read --allow-write --unstable src/mod.ts