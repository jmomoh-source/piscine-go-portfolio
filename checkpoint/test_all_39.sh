#!/bin/bash
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
cd /home/jmomoh/GO_PROJECTS_LEARN2EARN/piscine-go-portfolio/checkpoint

echo "Running full verification suite on all 39 exercises..."
passed=0
failed=0

for d in */; do
    d_clean="${d%/}"
    if [ "$d_clean" == "scratch" ]; then continue; fi
    
    echo -n "Testing $d_clean... "
    cd "$d_clean"
    out=$(go run . 2>&1)
    exit_code=$?
    cd ..
    
    if [ $exit_code -eq 0 ]; then
        echo "OK"
        ((passed++))
    else
        echo "FAILED: $out"
        ((failed++))
    fi
done

echo "----------------------------------------"
echo "Results: $passed Passed, $failed Failed out of 39."
