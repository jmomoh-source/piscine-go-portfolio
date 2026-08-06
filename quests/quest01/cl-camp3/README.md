# Quest01 — cl-camp3 (look)

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise focuses on using the `find` command to search files and directories with specific name patterns.
The task: create a file named `look` that contains the command to search in the current directory and its subfolders for:
- Everything that starts with `a`
- All files ending with `z`
- All files starting with `z` and ending with `a!`

## Instructions
- Use the `find` command
- Search recursively in the current directory (`.`)
- Apply multiple name patterns
- Output matching files and directories

## Implementation
`look`:

```bash
find . \( -name 'a*' -o -name '*z' -o -name 'z*a!' \)
```

## Explanation
- `find .` → start searching in the current directory
- `\( ... -o ... -o ... \)` → group multiple conditions with OR logic
- `-name 'a*'` → matches anything starting with `a`
- `-name '*z'` → matches anything ending with `z`
- `-name 'z*a!'` → matches anything starting with `z` and ending with `a!`

## Usage
Make the file executable:

```bash
chmod +x look
```

Run it:

```bash
./look
```

Expected output: a list of matching files and directories, e.g.

```text
./apple.txt
./amazing.doc
./buzz
./zebra-a!
```

## Skills Practiced
- Using `find` for recursive search
- Combining multiple conditions with `-o` (OR)
- Pattern matching with wildcards (`*`)
- Understanding file name filters

## Notes
- Remove `-type f` if you want to include directories in the results
- Always quote patterns to prevent shell expansion
- The exercise emphasizes reading `man find` to discover the right flags
