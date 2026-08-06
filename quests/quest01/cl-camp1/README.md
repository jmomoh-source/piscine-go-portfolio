# Quest01 — cl-camp1 (mastertheLS)

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise focuses on mastering the `ls` command with advanced options.  
The task: write a single command in a file named `mastertheLS` that lists files and directories in the current directory with specific formatting rules.

## Instructions
The command must:
- list files and directories of the current directory
- ignore hidden files, `.` and `..`
- separate results with commas only
- order results by access time, newest first
- append `/` to directory names

## Expected Output
Running the script should produce a single line like:

```text
file1,file2,dir1/,dir2/
```

## Files to Submit
- `mastertheLS`

## Implementation
`mastertheLS`:

```bash
#!/bin/bash
ls -ltu --time-style=+%s --hide='.*' --indicator-style=slash | awk '{$1=$2=$3=$4=$5=$6=$7=$8=""; sub(/^ +/, ""); print}' | paste -sd, -
```

## Explanation
- `ls -ltu` → long listing sorted by access time, newest first
- `--time-style=+%s` → use a stable timestamp format for parsing
- `--hide='.*'` → ignore hidden files
- `--indicator-style=slash` → append `/` to directory names
- `awk '{...}'` → remove the first eight `ls -l` columns and keep only the file names
- `paste -sd, -` → join all output lines into one comma-separated string

## Usage
Make the file executable:

```bash
chmod +x mastertheLS
```

Run it:

```bash
./mastertheLS
```

## Notes
- On macOS, GNU Core Utilities may be required; use `gls` instead of `ls`
- The exercise is all about reading `man ls` and finding the right flags
