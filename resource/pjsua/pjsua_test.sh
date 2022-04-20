#!/bin/bash

testpaths=`ls|grep test`

for p in $testpaths
do
    /usr/bin/screen -dmS pjsua pjsua --config-file=$p/config.cfg
done

