#!/bin/bash
while true
do
    netstat -an|grep 8080|grep EST -c
    sleep 4
done

