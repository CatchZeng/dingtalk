#!/bin/bash
shell_dir=$(dirname $0)
cd ${shell_dir}

cd ..

rm -rf test/mocks
mkdir -p test/mocks/message

# open with vscode
if which mockgen >/dev/null; then
  echo "mockgen has installed in PATH"
else
  echo "warning: 'mockgen' command has not installed in PATH"
  GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.3
fi

mockgen -package=mock_message -source=message.go > test/mocks/message/message.go