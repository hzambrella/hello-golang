#!/usr/bin/expect
git status
git add .
git status
git commit -m"1"
git status
set timeout 30
git push origin master





