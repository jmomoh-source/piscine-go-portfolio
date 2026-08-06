#!/bin/bash
cd /home/jmomoh/GO_PROJECTS_LEARN2EARN/piscine-go-portfolio
git checkout --orphan temp_branch
git add -A
git commit -m "GO-Lang-Foundation"
git branch -D main
git branch -m main
git push -f origin main
