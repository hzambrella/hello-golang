#!/bin/sh
#!/usr/bin/expect
git status
git add .
git status
git commit -m"1"
git status
set timeout 30
spawn git push origin master
expect "Username for 'https://github.com':"
send "hzambrella\r"
interact





