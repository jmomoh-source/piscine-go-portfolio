# Quest01 — File Permissions & Archiving

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise is about creating a specific filesystem layout with exact names, permissions, timestamps, and a symbolic link so that a custom `ls` command outputs the expected listing. It is a practical test of Unix permissions, timestamps, file types, and tar archiving.

## Skills Practiced
- Unix file permissions with `chmod`
- File timestamps with `touch -t`
- Symbolic links with `ln -s`
- Archiving with `tar`
- Output filtering with `ls`, `awk`, and `sed`

## Goal
Create the following entries:
- Files: `0`–`9`
- Directory: `A`
- Symbolic link: `3` pointing to `0`

Package only `done.tar` for submission.

## Expected Output
```bash
TZ=utc ls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'
```

Should produce:

```text
dr-------x 1986-01-05 00:00 0
-r------w- 1986-11-13 00:01 1
-rw----r-- 1988-03-05 00:10 2
lrwxrwxrwx 1990-02-16 00:11 3 -> 0
-r-x--x--- 1990-10-07 01:00 4
-r--rw---- 1990-11-07 01:01 5
-r--rw---- 1991-02-08 01:10 6
-r-x--x--- 1991-03-08 01:11 7
-rw----r-- 1994-05-20 10:00 8
-r------w- 1994-06-10 10:01 9
dr-------x 1995-04-10 10:10 A
```

## Quick Reference Table

| Entry | Permissions | Octal | Timestamp (UTC) |
|------|-------------|-------|-----------------|
| `0` | `dr-------x` | `401` | `1986-01-05 00:00` |
| `1` | `-r------w-` | `402` | `1986-11-13 00:01` |
| `2` | `-rw----r--` | `604` | `1988-03-05 00:10` |
| `3` | `lrwxrwxrwx` | link | `1990-02-16 00:11` |
| `4` | `-r-x--x---` | `510` | `1990-10-07 01:00` |
| `5` | `-r--rw----` | `460` | `1990-11-07 01:01` |
| `6` | `-r--rw----` | `460` | `1991-02-08 01:10` |
| `7` | `-r-x--x---` | `510` | `1991-03-08 01:11` |
| `8` | `-rw----r--` | `604` | `1994-05-20 10:00` |
| `9` | `-r------w-` | `402` | `1994-06-10 10:01` |
| `A` | `dr-------x` | `401` | `1995-04-10 10:10` |

## How to Reproduce

```bash
# Create files and directory
mkdir A
touch 0 1 2 3 4 5 6 7 8 9
ln -s 0 3

# Set permissions
chmod 401 A
chmod 402 1
chmod 604 2
chmod 510 4
chmod 460 5
chmod 460 6
chmod 510 7
chmod 604 8
chmod 402 9

# Set timestamps (UTC)
touch -t 198601050000 0
touch -t 198611130001 1
touch -t 198803050010 2
touch -h -t 199002160011 3
touch -t 199010070100 4
touch -t 199011070101 5
touch -t 199102080110 6
touch -t 199103080111 7
touch -t 199405201000 8
touch -t 199406101001 9
touch -t 199504101010 A
```

> Note: `touch -h` updates the symlink timestamp on GNU systems. On some systems, link timestamps may behave differently.

## Create Submission Archive

```bash
tar -cf done.tar *
ls
```

Only `done.tar` should be submitted.

## Notes for macOS Users
macOS does not ship with GNU `ls` by default. Install GNU Core Utilities and use `gls` instead:

```bash
TZ=utc gls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'
```

Symbolic link permissions may appear as `lrwxr-xr-x` depending on your OS, and that is acceptable for this exercise.
