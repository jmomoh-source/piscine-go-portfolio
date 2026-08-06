#!/bin/bash

if [[ -z "$HERO_ID" ]]; then
  echo "HERO_ID is not set" >&2
  exit 1
fi

curl -s https://acad.learn2earn.ng/assets/superhero/all.json \
  | jq -r ".[] | select(.id == ${HERO_ID}) | .connections.relatives"
