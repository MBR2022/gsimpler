#!/usr/bin/env bash

set -e
echo "" > coverage.txt
baseDir="github.com/MBR2022/gsimpler"
for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=tmp.out -covermode=atomic $d
    if [ -f tmp.out ]; then
        echo $d
        cat tmp.out >> coverage.txt
        targetDir=".${d#$baseDir}"
        cp tmp.out $targetDir/coverage.out
        rm tmp.out
    fi
done