# dingtalk

![Go](https://github.com/CatchZeng/dingtalk/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/CatchZeng/dingtalk/branch/master/graph/badge.svg)](https://codecov.io/gh/CatchZeng/dingtalk)
[![Go Report Card](https://goreportcard.com/badge/github.com/CatchZeng/dingtalk)](https://goreportcard.com/report/github.com/CatchZeng/dingtalk)
[![Release](https://img.shields.io/github/release/CatchZeng/dingtalk.svg)](https://github.com/CatchZeng/dingtalk/releases)
[![GoDoc](https://godoc.org/github.com/CatchZeng/dingtalk?status.svg)](https://pkg.go.dev/github.com/CatchZeng/dingtalk?tab=doc)

[中文](https://github.com/CatchZeng/dingtalk/blob/master/README.md)

> DingTalk (dingding) is the go implementation of the DingTalk robot. Support **Docker, Jenkinsfile,command line mode, module mode**, **signature security settings, chain syntax** to create messages, support **text, link, markdown、ActionCard、FeedCard** message types.

## Doc

[ding-doc](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)

## Feature

- [x] Support [Docker](https://github.com/CatchZeng/dingtalk#Docker)

- [x] Support [Jenkinsfile](https://github.com/CatchZeng/dingtalk#Jenkinsfile)

- [x] Support [module](https://github.com/CatchZeng/dingtalk/blob/master/READMEEN.md#use-as-module)

- [x] Support [Command Line Mode](https://github.com/CatchZeng/dingtalk/blob/master/READMEEN.md#use-as-command-line-tool)

- [x] Support [config.yaml](https://github.com/CatchZeng/dingtalk/blob/master/READMEEN.md#config.yaml)

- [x] Support sign

  <img src="https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1572261283991-f8e35f4d-6997-4a02-9704-843ee8f97464.png" width = 50% />

- [x] Text message

  <img src="https://img.alicdn.com/tfs/TB1jFpqaRxRMKJjy0FdXXaifFXa-497-133.png" width = 50% />

- [x] Link message

  <img src="https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1570679827267-6243216b-d1c3-48b7-9b1e-0f0b4211b50b.png" width = 50% />

- [x] Markdown message

  <img src="https://img.alicdn.com/tfs/TB1yL3taUgQMeJjy0FeXXXOEVXa-492-380.png" width = 50% />

- [x] ActionCard message

  <img src="https://img.alicdn.com/tfs/TB1nhWCiBfH8KJjy1XbXXbLdXXa-547-379.png" width = 50% />

  <img src="https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1570679939723-c1fb7861-5bcb-4c30-9e1b-033932f6b72f.png" width = 50% />

- [x] FeedCard message

  <img src="http://img01.taobaocdn.com/top/i1/LB1R2evQVXXXXXDapXXXXXXXXXX" width = 50% />

## Install

## with Docker

```shell
docker pull catchzeng/dingtalk
```

### binary

Go to [releases](https://github.com/CatchZeng/dingtalk/releases/) to download the binary executable file of the corresponding platform, and then add it to the PATH environment variable.

### with go get

```shell
go get github.com/CatchZeng/dingtalk
```

## Usage

### config.yaml

You can create `config.yaml` under `$/HOME/.dingtalk` and fill in the default values of `access_token` and `secret`.

```yaml
access_token: "1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f"
secret: "SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68"
```

### Docker

```shell
docker run catchzeng/dingtalk dingtalk text -t 1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f -s SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68 -c "docker test"
```

### Jenkinsfile

```shell
pipeline {
    agent {
        docker {
            image 'catchzeng/dingtalk:latest'
        }
    }
    environment {
        DING_TOKEN = '1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f'
        DING_SECRET = 'SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68'
    }
    stages {
        stage('notify') {
            steps {
                sh 'dingtalk link -t ${DING_TOKEN} -s ${DING_SECRET} -i "标题" -e "信息" -u "https://catchzeng.com/" -p "https://catchzeng.com/img/avatar-hux.jpg" -a'
            }
        }
    }
}
```

### Use as module

```go
package main

import (
    "log"

    "github.com/CatchZeng/dingtalk"
)

func main() {
	accessToken := "1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f"
    secret := "SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68"
    client := dingtalk.NewClient(accessToken, secret)

    msg := dingtalk.NewTextMessage().SetContent("测试文本&at 某个人").SetAt([]string{"177010xxx60"}, false)
    client.Send(msg)
}
```

### Use as command line tool

#### Demo

```shell
dingtalk text -t 1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f -s SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68 -c "测试命令行 & at 某个人" -m "177010xxx60","177010xxx61"
```

#### Help

```shell
$ dingtalk -h
dingtalk is a command line tool for DingTalk

Usage:
  dingtalk [command]

Available Commands:
  actionCard  send actionCard message with DingTalk robot
  feedCard    send feedCard message with DingTalk robot
  help        Help about any command
  link        send link message with DingTalk robot
  markdown    send markdown message with DingTalk robot
  text        send text message with DingTalk robot
  version     dingtalk version

Flags:
  -m, --atMobiles strings   atMobiles
  -h, --help                help for dingtalk
  -a, --isAtAll             isAtAll
  -s, --secret string       secret
  -t, --token string        access_token

Use "dingtalk [command] --help" for more information about a command.
```

## Stargazers

[![Stargazers over time](https://starchart.cc/CatchZeng/dingtalk.svg)](https://starchart.cc/CatchZeng/dingtalk)
