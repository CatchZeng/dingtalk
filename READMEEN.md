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

- [x] Support [environment variables](https://github.com/CatchZeng/dingtalk#environment%20variables)

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

### with go install

```sh
# Go 1.16+
go install github.com/CatchZeng/dingtalk@v1.5.0

# Go version < 1.16
go get -u github.com/CatchZeng/dingtalk@v1.5.0
```

## Usage

### config.yaml

You can create `config.yaml` under `$/HOME/.dingtalk` and fill in the default values of `access_token` and `secret`.

```yaml
access_token: "1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f"
secret: "SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68"
```

### environment variables

```sh
$ export ACCESS_TOKEN=1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f
$ export SECRET=SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68
$ dingtalk link -i "标题" -e "信息" -u "https://makeoptim.com/" -p "https://makeoptim.com/assets/img/logo.png" -a
```

You can also set a **prefix** for environment variables.

```sh
$ export DINGTALK_ENV_PREFIX="DINGTALK_"
$ export DINGTALK_ACCESS_TOKEN="1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f"
$ export DINGTALK_SECRET="SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68"
$ dingtalk link -i "标题" -e "信息" -u "https://makeoptim.com/" -p "https://makeoptim.com/assets/img/logo.png" -a
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
                sh 'dingtalk link -t ${DING_TOKEN} -s ${DING_SECRET} -i "标题" -e "信息" -u "https://makeoptim.com/" -p "https://makeoptim.com/assets/img/logo.png" -a'
            }
        }
    }
}
```

### Use as module

```sh
go get github.com/CatchZeng/dingtalk
```

```go
package main

import (
    "log"

    "github.com/CatchZeng/dingtalk/pkg/dingtalk"
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

```shell
$ dingtalk markdown -D -i "杭州天气" -e '## 杭州天气 @150XXXXXXXX
 > 9度，西北风1级，空气良89，相对温度73%
 > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)
 > ###### 10点20分发布 [天气](https://www.dingtalk.com)' -t 1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f -s SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68

{"msgtype":"markdown","markdown":{"title":"杭州天气","text":"## 杭州天气 @150XXXXXXXX\n \u003e 9度，西北风1级，空气良89，相对温度73%\n \u003e ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n \u003e ###### 10点20分发布 [天气](https://www.dingtalk.com)"},"at":{"atMobiles":[],"isAtAll":false}}
```

> -D: print the message content

#### Help

```shell
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
  -t, --access_token string   access_token
  -m, --atMobiles strings     atMobiles
  -D, --debug                 debug
  -h, --help                  help for dingtalk
  -a, --isAtAll               isAtAll
  -s, --secret string         secret

Use "dingtalk [command] --help" for more information about a command.
```

## Stargazers

[![Stargazers over time](https://starchart.cc/CatchZeng/dingtalk.svg)](https://starchart.cc/CatchZeng/dingtalk)
