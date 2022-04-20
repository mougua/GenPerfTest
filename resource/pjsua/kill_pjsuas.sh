#!/bin/bash

pids=`ps aux|grep pjsua|grep -v kill|grep -v grep|grep -v vim|awk '{print $2}'`

for pid in $pids
do
    kill $pid 
done

