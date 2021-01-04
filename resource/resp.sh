#!/bin/bash
# echo $1 >> req.txt

os=$(uname)
who=$(whoami)
resp="$os|$who"
echo -n $resp
