#!/bin/bash

PROJECT=$1

if [ -z "$PROJECT" ]; then
    PROJECT=quantum
fi

SRTPATH=$(cd "$(dirname "$0")"; pwd)

cd $SRTPATH/../src/$PROJECT

files=`find . -name \*.go | grep -v ffjson`

total=0
for file in $files; do
    lines=`wc -l $file | awk '{print $1}'`
    ((total+=lines))
done
echo "Total:$total"
