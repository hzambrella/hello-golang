#!/bin/sh
git status
git add .
git status
git commit -m"1"
git status
set timeout 30
spawn git push origin master
./login.sh




