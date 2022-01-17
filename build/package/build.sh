#!/bin/bash
shell_dir=$(dirname $0)
cd ${shell_dir}

# check params
if [[ ! $1 ]]; then
    echo "image tag is null"; exit 1;
else
    echo "image tag: $1"
fi

cd ../../

docker build -t $1 -f build/package/Dockerfile .
