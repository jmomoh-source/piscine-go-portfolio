# Quest01 — cl-camp2 (r)

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise is about working with plain text files and understanding how the `cat` command displays content.
The task: create a file named `r` that prints the letter **R** followed by a newline when read with `cat`.

## Instructions
- Create a file named `r`
- The file must contain exactly one line: `R`
- A line ends with the newline character (`\n`)
- When tested with `cat -e`, the output must be:

```bash
$ cat -e r
R$
$
```

## Implementation
To create the file:

```bash
echo "R" > r
```

This writes `R` followed by a newline into the file.

Verify:

```bash
cat -e r
```

Expected output:

```text
R$
$
```

## Explanation
- `echo "R"` prints `R` followed by a newline
- `>` redirects the output into the file `r`
- `cat -e r` shows file contents, with `$` marking the end of each line

## Skills Practiced
- Creating files with text content
- Understanding newline characters (`\n`)
- Using `cat` with the `-e` option to visualize line endings
- Basic shell redirection (`>`)

## Notes
- Do not add extra spaces or lines — the file must contain only `R` and a newline
- On different systems, `cat -e` may behave slightly differently, but the `$` marker always indicates the end of a line
