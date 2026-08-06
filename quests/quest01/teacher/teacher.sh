#!/bin/bash

# Step 1: locate the interview file (consistent command across mystery folders)
FILEPATH=$(find mystery -type f -path '*/interviews/*' -print | head -n 1)

if [[ -z "$FILEPATH" ]]; then
  echo "" >&2
  exit 1
fi

# Extract interview number (basename without extension) and export it
INTERVIEW=$(basename "$FILEPATH" | cut -d'.' -f1)
export INTERVIEW

# Step 2: print the interview number
echo "$INTERVIEW"

# Step 3: print the content of the interview file
cat "$FILEPATH"

# Step 4: print the MAIN_SUSPECT environment variable (if set)
echo "${MAIN_SUSPECT:-}"
