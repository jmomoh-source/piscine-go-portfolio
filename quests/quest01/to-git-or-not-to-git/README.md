# Quest01 — To Git or Not to Git

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise builds on the previous quest by extending Bash scripting and JSON parsing.  
The task: create a script `to-git-or-not-to-git.sh` that fetches superhero data and prints the **name, power, and gender** of the hero with `id: 170`.

## Instructions
- Fetch data from: https://acad.learn2earn.ng/assets/superhero/all.json  
- Use `curl` to download the JSON  
- Use `jq` to parse and filter the data  
- Output must match exactly:

```bash
$ bash to-git-or-not-to-git.sh
Chameleon
28
Male
$
```

## Implementation
`to-git-or-not-to-git.sh`:

```bash
#!/bin/bash

curl -s https://acad.learn2earn.ng/assets/superhero/all.json \
  | jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'
```

## Explanation
- `#!/bin/bash` → Shebang line, tells the system to run the script with Bash
- `curl -s` → Fetches the JSON silently (no progress bar)
- `jq -r` → Processes JSON and outputs raw text (no quotes)
- `.[]` → Iterates through each superhero object
- `select(.id == 170)` → Filters the hero with id 170
- `.name, .powerstats.power, .appearance.gender` → Extracts three fields, each printed on a new line

## Usage
Make the script executable:

```bash
chmod +x to-git-or-not-to-git.sh
```

Run it:

```bash
./to-git-or-not-to-git.sh
```

Expected output:

```text
Chameleon
28
Male
```

## Skills Practiced
- Bash scripting with multiple commands
- API calls with curl
- JSON parsing with jq
- Nested field access in JSON (`.powerstats.power`, `.appearance.gender`)
- Raw output formatting with `-r`

## Notes
- Ensure `jq` is installed (`sudo apt install jq` or `brew install jq`)
- The `-r` flag is critical to remove quotes from the output
- The comma operator in jq prints multiple fields on separate lines
