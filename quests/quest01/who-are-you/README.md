# Quest01 — Who Are You?

![Shell](https://img.shields.io/badge/shell-bash-green)
![License](https://img.shields.io/badge/license-MIT-blue)

## Overview
This exercise introduces Bash scripting, API calls, and JSON parsing.  
The scenario: *"You just woke up in a dark alley... You can not remember who you are... The only thought that comes to your mind is a tag that says: subject Id: 70."*  
The task is to create a script `who-are-you.sh` that fetches superhero data and prints the name of the hero with `id: 70`.

## Instructions
- Fetch data from: https://acad.learn2earn.ng/assets/superhero/all.json  
- Use `curl` to download the JSON  
- Use `jq` to parse and filter the data  
- Output must match exactly:

```bash
./who-are-you.sh | cat -e
"name"$
```

## Files to Submit
- `who-are-you.sh`

## Implementation
Create `who-are-you.sh` with the following content:

```bash
#!/bin/bash

curl -s https://acad.learn2earn.ng/assets/superhero/all.json \
  | jq '.[] | select(.id == 70) | .name'
```

## Explanation
- `#!/bin/bash` → Shebang line tells the system to run the script with Bash
- `curl -s` → Fetches the JSON silently (no progress bar)
- `jq '.[] | select(.id == 70) | .name'` → Iterates through the array, filters the object with `id == 70`, and extracts the `name` field

## Usage
Make the script executable:

```bash
chmod +x who-are-you.sh
```

Run it:

```bash
./who-are-you.sh | cat -e
```

Expected output:

```text
"Batman"$
```

(The actual name depends on the dataset — `id: 70` in the JSON.)

## Skills Practiced
- Bash scripting basics (`#!/bin/bash`)
- API calls with `curl`
- JSON parsing with `jq`
- Filtering arrays and extracting fields
- Command chaining with pipes

## Notes
- Ensure `jq` is installed (`sudo apt install jq` or `brew install jq`)
- Output must include quotes and end with `$` when piped through `cat -e`
