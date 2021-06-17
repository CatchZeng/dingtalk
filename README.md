# dingtalk

![Go](https://github.com/CatchZeng/dingtalk/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/CatchZeng/dingtalk/branch/master/graph/badge.svg)](https://codecov.io/gh/CatchZeng/dingtalk)
[![Go Report Card](https://goreportcard.com/badge/github.com/CatchZeng/dingtalk)](https://goreportcard.com/report/github.com/CatchZeng/dingtalk)
[![Release](https://img.shields.io/github/release/CatchZeng/dingtalk.svg)](https://github.com/CatchZeng/dingtalk/releases)
[![GoDoc](https://godoc.org/github.com/CatchZeng/dingtalk?status.svg)](https://pkg.go.dev/github.com/CatchZeng/dingtalk?tab=doc)

[English](https://github.com/CatchZeng/dingtalk/blob/master/READMEEN.md)

> DingTalk(dingding) 是钉钉机器人的 go 实现。支持 **Docker、Jenkinsfile、命令行模式，module 模式**；支持**加签**安全设置，支持**链式语法**创建消息；支持**文本、链接、Markdown、ActionCard、FeedCard** 消息类型。

> 注：使用飞书的小伙伴，可以使用[飞书（feishu）版](https://github.com/CatchZeng/feishu)。

## 文档

[钉钉文档](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)

## 特性

- [x] 支持[Docker](https://github.com/CatchZeng/dingtalk#Docker)

- [x] 支持[Jenkinsfile](https://github.com/CatchZeng/dingtalk#Jenkinsfile)

- [x] 支持[module](https://github.com/CatchZeng/dingtalk#%E4%BD%9C%E4%B8%BA-module)

- [x] 支持[命令行模式](https://github.com/CatchZeng/dingtalk#%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%B7%A5%E5%85%B7)

- [x] 支持[配置文件](https://github.com/CatchZeng/dingtalk#%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6)

- [x] 支持加签

  <img src="https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1572261283991-f8e35f4d-6997-4a02-9704-843ee8f97464.png" width = 50% />

- [x] Text 消息

  <img src="https://img.alicdn.com/tfs/TB1jFpqaRxRMKJjy0FdXXaifFXa-497-133.png" width = 50% />

- [x] Link 消息

  <img src="https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1570679827267-6243216b-d1c3-48b7-9b1e-0f0b4211b50b.png" width = 50% />

- [x] Markdown 消息

  <img src="https://img.alicdn.com/tfs/TB1yL3taUgQMeJjy0FeXXXOEVXa-492-380.png" width = 50% />

- [x] ActionCard 消息

    <img src="https://img.alicdn.com/tfs/TB1nhWCiBfH8KJjy1XbXXbLdXXa-547-379.png" width = 50% />

    <img src="https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1570679939723-c1fb7861-5bcb-4c30-9e1b-033932f6b72f.png" width = 50% />

- [x] FeedCard 消息

  <img src="http://img01.taobaocdn.com/top/i1/LB1R2evQVXXXXXDapXXXXXXXXXX" width = 50% />

## 安装

## Docker 安装

```shell
docker pull catchzeng/dingtalk
```

### 二进制安装

到 [releases](https://github.com/CatchZeng/dingtalk/releases/) 下载相应平台的二进制可执行文件，然后加入到 PATH 环境变量即可。

### go get 安装

```shell
go get github.com/CatchZeng/dingtalk
```

## 使用方法

### 配置文件

可以在 `$/HOME/.dingtalk` 下创建 `config.yaml` 填入 `access_token` 和 `secret` 默认值。

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

### 作为 module

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

### 命令行工具

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
