#!/bin/sh
#sudo aptitude install expect
#cp push_example.sh push.sh
#if spawn is not found 
#don't use sh push.sh
#use chmod 755 push.sh or chmod a-x push.sh
#the ./push.sh

read -p "input commit measure:"  val
echo $val
git status
git add .
git status
git commit -m"$val"
git status
#read -p "are you sure to commit ?[Y/N]"val2
set timeout 30
expect <<!
spawn git push origin master
expect "Username"
send "your name\r"
expect "Password"
send "your password\r"
expect eof
!





