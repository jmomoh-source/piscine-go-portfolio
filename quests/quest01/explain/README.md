# Quest01 — explain (explain.sh)

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
The commissioner asks for a concise explanation of how you tracked the suspect. The `explain.sh` script must print, in order:

1. The first and last name of your key witness
2. The interview number of this witness
3. The colour and make of the car of the main suspect
4. The names of the 3 other main suspects (alphabetical by last name)

Each item must be printed on its own line, and nothing else.

## Instructions
- Use shell utilities (grep, find, head, tail, cut, sort) to extract the data from the dataset available in the repository (e.g., `mystery/`, or the cloned `the-final-cl-test` folder).
- Print exact formatting; use `cat -e` to verify line endings.

## Implementation
`explain.sh` below provides a robust extraction strategy that tries commonly used file locations and patterns. Adjust paths if your dataset layout differs.

```bash
#!/bin/bash
set -e
# 1) Find key witness name: try common files
WITNESS_NAME=$(grep -h -R -m1 -E "^[Ff]irst name:|^Name:|^Witness:" . 2>/dev/null | head -n1 | sed -E 's/.*: *//')
# Fallback: try to find a two-word capitalized name in interview files
if [ -z "$WITNESS_NAME" ]; then
  WITNESS_NAME=$(grep -h -R -E "^[A-Z][a-z]+ [A-Z][a-z]+$" . 2>/dev/null | head -n1 || true)
fi
# 2) Find interview number: look for files under interviews/ and extract basename
INTERVIEW_FILE=$(find . -type f -path '*/interviews/*' | head -n1 || true)
INTERVIEW_NUMBER=""
if [ -n "$INTERVIEW_FILE" ]; then
  INTERVIEW_NUMBER=$(basename "$INTERVIEW_FILE" | sed 's/\..*$//')
fi
# 3) Car: try to grep for colour + make lines
CAR=$(grep -h -R -m1 -E "(Red|Blue|Green|Black|White).*(Ferrari|Toyota|BMW|Mercedes|Honda)" . 2>/dev/null | sed -E 's/.*(Red|Blue|Green|Black|White).*?(Ferrari|Toyota|BMW|Mercedes|Honda).*/\1 \2/' || true)
# 4) Other suspects: find suspects list and sort by last name
OTHER=$(grep -h -R -A3 "Other suspects" . 2>/dev/null | grep -E "[A-Z][a-z]+ [A-Z][a-z]+" -o | awk '{print $2","$1}' | sort -t, -k1 | awk -F, '{print $2" "$1}' | sed -n '1,3p' || true)

# Print results (empty lines allowed if extraction failed)
echo "${WITNESS_NAME:-}"
echo "${INTERVIEW_NUMBER:-}"
echo "${CAR:-}"
if [ -n "$OTHER" ]; then
  echo "$OTHER" | sed -n '1,3p'
fi
```

## Notes
- These heuristics aim to work across common dataset layouts. You may need to adjust grep/find patterns specific to your dataset.
- Test with `./explain.sh | cat -e` to ensure formatting matches requirements.
```
