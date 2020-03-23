#!/bin/bash
##################################################################
# Copyright(C) 2020-2020. All right reserved.
# 
# Filename: run.sh
# Author: ahaoozhang
# Date: 2020-03-23 22:45:26 (Monday)
# Describe: 
##################################################################

if [ ! -d "./log" ]; then
mkdir ./log
fi

if [ ! -f "./log/debug.log" ]; then
touch ./log/debug.log
fi

nohup ./bin/GradeManager > ./log/debug.log 2>&1 &!
