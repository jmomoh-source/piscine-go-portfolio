# Quest01 — teacher (teacher.sh)

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise builds a reproducible training script that documents how the mystery was solved.
The task: create a `teacher.sh` script that isolates the key interview number into an environment variable, prints it, shows the interview content, and prints the `MAIN_SUSPECT` environment variable.

## Instructions
The script must perform the following steps:
1. Isolate the interview number into an environment variable using the same file-discovery command used during the investigation.
2. Print the environment variable (the interview number only).
3. Print the content of the interview file.
4. Print the value of the `MAIN_SUSPECT` environment variable.

## Implementation
`teacher.sh`:

```bash
#!/bin/bash

# Find the interview file (same method across mystery folders)
FILEPATH=$(find mystery -type f -path '*/interviews/*' -print | head -n 1)

if [[ -z "$FILEPATH" ]]; then
  echo "No interview file found" >&2
  exit 1
fi

# Extract interview number (basename without extension), export for downstream use
INTERVIEW=$(basename "$FILEPATH" | cut -d'.' -f1)
export INTERVIEW

# Step 2: print the interview number
echo "$INTERVIEW"

# Step 3: show the interview content
cat "$FILEPATH"

# Step 4: print the MAIN_SUSPECT environment variable (may be set externally)
echo "${MAIN_SUSPECT:-}"
```

## Usage
Make the script executable:

```bash
chmod +x teacher.sh
```

Run it inside a mystery folder (make sure `MAIN_SUSPECT` is exported if required):

```bash
./teacher.sh
```

Expected output (example):

```text
123456
This is the interview content...
John Doe
```

## Notes
- The `find` command targets files under `mystery/*/interviews/` and selects the first match; this selection method is intentionally consistent across mystery folders.
- Ensure no extra spaces or blank lines are printed beyond what's requested.
- `MAIN_SUSPECT` is expected to be set in the environment by the mystery dataset or a wrapper script.
