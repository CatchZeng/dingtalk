#!/bin/bash
shell_dir=$(dirname $0)
cd ${shell_dir}

cd ..

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -gcflags=-l -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
