#!/bin/sh
git status
git add .
git status
git commit -m"1"
git status
set timeout 30
expect <<!
spawn git push origin master
echo login.sh
expect "Username"
send "hzambrella/r"
expect ">"
expect eof
!




