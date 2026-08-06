#!/bin/bash
set -e
# 1) Try to extract witness name from common labelled lines
WITNESS_NAME=$(grep -h -R -m1 -E "^[Ff]irst name:|^Name:|^Witness:" . 2>/dev/null | head -n1 | sed -E 's/.*: *//')
if [ -z "$WITNESS_NAME" ]; then
  WITNESS_NAME=$(grep -h -R -E "^[A-Z][a-z]+ [A-Z][a-z]+$" . 2>/dev/null | head -n1 || true)
fi

# 2) Interview number: prefer filename under interviews/
INTERVIEW_FILE=$(find . -type f -path '*/interviews/*' | head -n1 || true)
INTERVIEW_NUMBER=""
if [ -n "$INTERVIEW_FILE" ]; then
  INTERVIEW_NUMBER=$(basename "$INTERVIEW_FILE" | sed 's/\..*$//')
else
  INTERVIEW_NUMBER=$(grep -h -R -m1 -E "Interview[ _]?[Nn]o\.?|Interview[ _]?[0-9]+" . 2>/dev/null | sed -E 's/.*([0-9]{3,}).*/\1/' || true)
fi

# 3) Car colour and make: look for common patterns
CAR=$(grep -h -R -m1 -E "(Red|Blue|Green|Black|White).*(Ferrari|Toyota|BMW|Mercedes|Honda|Ford)" . 2>/dev/null | sed -E 's/.*(Red|Blue|Green|Black|White).*?(Ferrari|Toyota|BMW|Mercedes|Honda|Ford).*/\1 \2/' || true)

# 4) Other suspects: find suspects lists and output 3 names sorted by last name
OTHER=$(grep -h -R -E "[A-Z][a-z]+ [A-Z][a-z]+" . 2>/dev/null | sort -u | awk '{print $2","$1}' | sort -t, -k1 | awk -F, '{print $2" "$1}' | sed -n '1,3p' || true)

# Print results
echo "${WITNESS_NAME:-}"
echo "${INTERVIEW_NUMBER:-}"
echo "${CAR:-}"
if [ -n "$OTHER" ]; then
  echo "$OTHER" | sed -n '1,3p'
fi
