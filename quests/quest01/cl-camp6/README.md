# Quest01 — cl-camp6 (countfiles.sh)

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise focuses on counting files and directories recursively.
The task: create a script `countfiles.sh` that prints the **number** of regular files and folders contained in the current directory and all its subfolders (including the current directory itself).

## Instructions
- Search from the current directory and all subfolders
- Count both files and directories
- Print only the number (no extra text)
- Output must match exactly

## Implementation
`countfiles.sh`:

```bash
#!/bin/bash

find . | wc -l
```

## Explanation
- `find .` → recursively lists all files and directories starting from the current directory
- `wc -l` → counts the number of lines (each line corresponds to one file or directory)
- The result is just a number, with no extra formatting

## Usage
Make the script executable:

```bash
chmod +x countfiles.sh
```

Run it:

```bash
./countfiles.sh | cat -e
```

Expected output (example):

```text
12$
$
```

## Skills Practiced
- Recursive search with `find`
- Counting results with `wc -l`
- Understanding how directories and files are represented in listings
- Producing clean numeric output

## Notes
- The count includes the current directory (`.`) itself
- If you want to exclude `.` and `..`, you can refine the find command with `-mindepth 1`
- `cat -e` is used in the test to visualize the newline at the end of the output
