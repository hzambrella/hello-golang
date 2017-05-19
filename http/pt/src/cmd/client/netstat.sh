#!/bin/bash
while true
do
    netstat -an|grep 9600|grep EST -c
    sleep 4
done

