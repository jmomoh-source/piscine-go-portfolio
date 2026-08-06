# Quest01 — cl-camp7

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise is about handling unusual filenames and ensuring exact content.
The task: create a file named **"\?$*'ChouMi'*$?\"** that contains only `01` followed by a newline.

## Instructions
- The filename must be exactly: `"\?$*'ChouMi'*$?\"`
- The file content must be exactly: `01` (with a newline at the end)
- When listed with `ls | cat -e`, the output must show:

```bash
$ ls | cat -e
"\?$*'ChouMi'*$?\"$
$
```

## Implementation
To create the file:

```bash
echo "01" > "\?$*'ChouMi'*$?\"
```

Verify:

```bash
ls | cat -e
```

Expected output:

```text
"\?$*'ChouMi'*$?\"$
$
```

## Explanation
- `echo "01"` → writes `01` followed by a newline
- `>` → redirects the output into the file
- Quoting the filename ensures the shell interprets special characters literally
- `cat -e` → shows `$` at the end of each line, confirming the newline

## Skills Practiced
- Creating files with unusual names containing special characters
- Proper use of quotes in shell to handle special characters
- Understanding newline characters (`\n`)
- Using `cat -e` to visualize line endings

## Notes
- Always wrap the filename in quotes when creating or accessing it
- Avoid escaping characters manually; quoting the entire string is safer
- The file must contain only `01` and a newline — no extra spaces or lines
