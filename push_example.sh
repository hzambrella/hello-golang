#!/bin/sh
#sudo aptitude install expect
#cp push_example.sh push.sh
#if spawn is not found 
#don't use sh push.sh
#use chmod 755 push.sh or chmod a-x push.sh
#the ./push.sh

git status
git add .
git status
git commit -m"1"
git status
set timeout 30
expect <<!
spawn git push origin master
expect "Username"
send "your username\r"
expect "Password"
send "your password\r"
expect eof
!




