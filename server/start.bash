#!/bin/bash
# 先杀死之前的进程,然后判断执行环境

process=`pgrep just.com |cut -d : -f1`
if [ "$process"x != ""x ]
then
    kill -9 $process
fi

rm -f ./bin/just.com
go install just.com

env=dev
if  [ "$1"x = "production"x ]   # 这个x是故意加上去的
then
    env=production
fi

nohup ./bin/just.com $env > ./log/log 2>&1 &