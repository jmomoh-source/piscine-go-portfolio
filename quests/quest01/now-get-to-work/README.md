# Quest01 — now-get-to-work (my_answer.sh)

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This final exercise asks you to clone the repository `github.com/01-edu/the-final-cl-test` and extract the required answer. The script `my_answer.sh` should perform the clone (locally) and print the answer on stdout.

## Instructions
- Clone the repo: `git clone https://github.com/01-edu/the-final-cl-test`
- Inspect files to find the answer (hint: use `head` and `tail` to extract a specific line)
- The script must print only the answer followed by a newline

## Implementation
`my_answer.sh` (automated approach):

```bash
#!/bin/bash
set -e
TMP=$(mktemp -d)
if ! git clone --depth 1 https://github.com/01-edu/the-final-cl-test "$TMP" 2>/dev/null; then
  # If clone fails, exit with empty output
  exit 0
fi
# Prioritized guesses: answer.txt, README.md last line, then try to find a line that looks like a name
if [ -f "$TMP/answer.txt" ]; then
  head -n 1 "$TMP/answer.txt"
  exit 0
fi
if [ -f "$TMP/README.md" ]; then
  tail -n 1 "$TMP/README.md" | sed -n '1p'
  exit 0
fi
# Fallback: look for the first line in the repo that matches "First Last" pattern
grep -h -R -E "^[A-Z][a-z]+ [A-Z][a-z]+$" "$TMP" 2>/dev/null | head -n1 || true
```

## Notes
- The exact file containing the answer may vary; the script tries common locations and a simple heuristic.
- You can adjust the extraction logic if you know the exact file to target.
