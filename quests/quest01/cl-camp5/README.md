# Quest01 — cl-camp5 (lookagain.sh)

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise continues the practice of using `find` and text processing tools.
The task: create a script `lookagain.sh` that searches recursively for all files ending with `.sh`, strips the extension, and prints the names in **descending order**.

## Instructions
- Search from the current directory and all subfolders
- Match files ending with `.sh`
- Output only the base name without `.sh`
- Sort results in descending order
- Output must match exactly

## Implementation
`lookagain.sh`:

```bash
#!/bin/bash

find . -type f -name "*.sh" \
  | sed 's/\.sh$//' \
  | sort -r
```

## Explanation
- `find . -type f -name "*.sh"` → recursively find all files ending with `.sh`
- `sed 's/\.sh$//'` → remove the `.sh` extension from each filename
- `sort -r` → sort results in reverse (descending) order

## Usage
Make the script executable:

```bash
chmod +x lookagain.sh
```

Run it:

```bash
./lookagain.sh | cat -e
```

Expected output (example):

```text
file3$
file2$
file1$
```

## Skills Practiced
- Recursive search with `find`
- Pattern matching with wildcards
- Text transformation with `sed`
- Sorting output with `sort`
- Combining commands with pipes

## Notes
- Always quote the pattern `*.sh` to prevent shell expansion
- `sort -r` ensures descending order; without `-r` it would be ascending
- The `$` markers in `cat -e` confirm line endings
