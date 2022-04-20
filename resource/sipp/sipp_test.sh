#!/bin/bash

sipp -i 172.16.23.242 -p 7766 -sf uac.xml -inf t.csv 172.16.23.52:6060 -l 159 -rtp_echo -aa -r 300 -rp 1000 -m 20000000


