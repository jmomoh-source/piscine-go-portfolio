# Quest01 — cl-camp8 (skip.sh)

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise focuses on filtering command output using text processing tools.
The task: create a script `skip.sh` that runs `ls -l` and prints every other line, **skipping 1 line out of 2 starting with the first one**.

## Instructions
- Run `ls -l` to list files in long format
- Skip the first line, then print every second line
- Output must match exactly when tested with `cat -e`

## Implementation
`skip.sh`:

```bash
#!/bin/bash

ls -l | sed -n '2~2p'
```

## Explanation
- `ls -l` → lists files in long format
- `sed -n '2~2p'` → starts at line 2 and prints every 2nd line thereafter
- `2~2` means "begin at line 2, then every 2 lines"
- `p` means "print"
- This effectively skips line 1, prints line 2, skips line 3, prints line 4, and so on

## Usage
Make the script executable:

```bash
chmod +x skip.sh
```

Run it:

```bash
./skip.sh | cat -e
```

Expected output (example):

```text
-rw-r--r--  1 user group  123 May  5 14:00 file1$
-rw-r--r--  1 user group  456 May  5 14:01 file2$
$
```

## Skills Practiced
- Using `ls -l` for detailed file listings
- Filtering output with `sed`
- Understanding line addressing (`2~2`)
- Combining commands with pipes

## Notes
- `awk 'NR % 2 == 0'` is another valid approach (prints only even-numbered lines)
- `cat -e` is used in the test to visualize line endings (`$`)
- The exercise emphasizes reading `man sed` or `man awk` to discover useful options
