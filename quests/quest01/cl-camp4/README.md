# Quest01 — cl-camp4 (myfamily.sh)

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise extends Bash scripting and JSON parsing by introducing environment variables.
The task: create a script `myfamily.sh` that shows the **relatives** of a superhero, based on the ID stored in the environment variable `HERO_ID`.

## Instructions
- Fetch data from: https://acad.learn2earn.ng/assets/superhero/all.json
- Use `curl` to download the JSON
- Use `jq` to parse and filter the data
- The subject is chosen dynamically using the environment variable `HERO_ID`
- Output must remove quotes and match exactly

## Implementation
`myfamily.sh`:

```bash
#!/bin/bash

if [[ -z "$HERO_ID" ]]; then
  echo "HERO_ID is not set" >&2
  exit 1
fi

curl -s https://acad.learn2earn.ng/assets/superhero/all.json \
  | jq -r ".[] | select(.id == ${HERO_ID}) | .connections.relatives"
```

## Explanation
- `#!/bin/bash` → Shebang line, tells the system to run the script with Bash
- `if [[ -z "$HERO_ID" ]]` → checks that the environment variable is set
- `curl -s` → fetches the JSON silently (no progress bar)
- `jq -r` → processes JSON and outputs raw text (removes quotes)
- `.[]` → iterates through each superhero object
- `select(.id == ${HERO_ID})` → filters the hero whose `id` matches the environment variable
- `.connections.relatives` → extracts the relatives field

## Usage
Set the environment variable:

```bash
export HERO_ID=1
```

Run the script:

```bash
./myfamily.sh
```

Expected output:

```text
Marlo Chandler-Jones (wife); Polly (aunt); Mrs. Chandler (mother-in-law); Keith Chandler, Ray Chandler, three unidentified others (brothers-in-law); unidentified father (deceased); Jackie Shorr (alleged mother; unconfirmed)
```

## Skills Practiced
- Bash scripting with environment variables
- API calls with `curl`
- JSON parsing with `jq`
- Nested field access in JSON (`.connections.relatives`)
- Raw output formatting with `-r`

## Notes
- Ensure `jq` is installed (`sudo apt install jq` or `brew install jq`)
- Always export `HERO_ID` before running the script
- The relatives field varies depending on the chosen superhero ID
