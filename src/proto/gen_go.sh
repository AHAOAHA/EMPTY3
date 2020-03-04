#!/bin/bash
##################################################################
# Copyright(C) 2020-2020. All right reserved.
# 
# Filename: gen_go.sh
# Author: ahaoozhang
# Date: 2020-03-04 21:54:09 (Wednesday)
# Describe: 
##################################################################
protoc --go_out=. DataCenter.proto
