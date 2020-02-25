# dingtalk

[中文](https://github.com/CatchZeng/dingtalk/blob/master/README.md)

> DingTalk (dingding) is the go implementation of the DingTalk robot. Support `Docker`, `Jenkinsfile`,`command line` mode, `module` mode, `signature security` settings, `chain syntax` to create messages, support `text, link, markdown、ActionCard、FeedCard` message types.

## Doc

[ding-doc](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)

## Feature

- [x] Support [Docker](https://github.com/CatchZeng/dingtalk#Docker)

- [x] Support [Jenkinsfile](https://github.com/CatchZeng/dingtalk#Jenkinsfile)

- [x] Support [module](https://github.com/CatchZeng/dingtalk/blob/master/READMEEN.md#use-as-module)

- [x] Support [Command Line Mode](https://github.com/CatchZeng/dingtalk/blob/master/READMEEN.md#use-as-command-line-tool)

- [x] Support sign

  ![sign](https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1572261283991-f8e35f4d-6997-4a02-9704-843ee8f97464.png)

- [x] Text message

  ![text](https://img.alicdn.com/tfs/TB1jFpqaRxRMKJjy0FdXXaifFXa-497-133.png)

- [x] Link message

  ![link](https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1570679827267-6243216b-d1c3-48b7-9b1e-0f0b4211b50b.png)

- [x] Markdown message

  ![markdown](https://img.alicdn.com/tfs/TB1yL3taUgQMeJjy0FeXXXOEVXa-492-380.png)

- [x] ActionCard message

  ![ActionCard1](https://img.alicdn.com/tfs/TB1nhWCiBfH8KJjy1XbXXbLdXXa-547-379.png)

  ![ActionCard2](https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1570679939723-c1fb7861-5bcb-4c30-9e1b-033932f6b72f.png)

- [x] FeedCard message

  ![ActionCard1](http://img01.taobaocdn.com/top/i1/LB1R2evQVXXXXXDapXXXXXXXXXX)

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

    "github.com/CatchZeng/dingtalk/client"
    "github.com/CatchZeng/dingtalk/message"
)

func main() {
    dingTalk := client.DingTalk{
        AccessToken: "1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f",
        Secret:      "SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68",
    }

    msg := message.NewTextMessage().SetContent("测试文本&at 某个人").SetAt([]string{"177010xxx60"}, false)
    dingTalk.Send(msg)
}
```

### Use as command line tool

#### Demo

```shell
dingtalk text -t 1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f -s SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68 -c "测试命令行 & at 某个人" -m "177010xxx60","177010xxx61"
```

#### Help

- dingtalk

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

- text

  ```shell
  $ dingtalk text -h
  send text message with DingTalk robot

  Usage:
    dingtalk text [flags]

  Flags:
    -c, --content string   content
    -h, --help             help for text

  Global Flags:
    -m, --atMobiles strings       atMobiles
    -a, --isAtAll                 isAtAll
    -s, --secret string           secret
    -t, --token string            access_token
  ```

- link

  ```shell
  $ dingtalk link -h
  send link message with DingTalk robot

  Usage:
    dingtalk link [flags]

  Flags:
    -h, --help                help for link
    -u, --messageURL string   messageURL
    -p, --picURL string       picURL
    -e, --text string         text
    -i, --title string        title

  Global Flags:
    -m, --atMobiles strings       atMobiles
    -a, --isAtAll                 isAtAll
    -s, --secret string           secret
    -t, --token string            access_token
  ```

- markdown

  ```shell
  $ dingtalk markdown -h
  send markdown message with DingTalk robot

  Usage:
    dingtalk markdown [flags]

  Flags:
    -h, --help           help for markdown
    -e, --text string    text
    -i, --title string   title

  Global Flags:
    -m, --atMobiles strings       atMobiles
    -a, --isAtAll                 isAtAll
    -s, --secret string           secret
    -t, --token string            access_token
  ```

- actionCard

  ```shell
  $ dingtalk actionCard -h
  send actionCard message with DingTalk robot

  Usage:
    dingtalk actionCard [flags]

  Flags:
    -c, --btnActionURLs strings   btnActionURLs
    -o, --btnOrientation string   btnOrientation
    -b, --btnTitles strings       btnTitles
    -h, --help                    help for actionCard
    -d, --hideAvatar string       hideAvatar
    -n, --singleTitle string      singleTitle
    -u, --singleURL string        singleURL
    -e, --text string             text
    -i, --title string            title

  Global Flags:
    -m, --atMobiles strings   atMobiles
    -a, --isAtAll             isAtAll
    -s, --secret string       secret
    -t, --token string        access_token
  ```

- feedCard

  ```shell
  dingtalk feedCard -h
  send feedCard message with DingTalk robot

  Usage:
    dingtalk feedCard [flags]

  Flags:
    -h, --help                  help for feedCard
    -u, --messageURLs strings   messageURLs
    -p, --picURLs strings       picURLs
    -i, --titles strings        titles

  Global Flags:
    -m, --atMobiles strings   atMobiles
    -a, --isAtAll             isAtAll
    -s, --secret string       secret
    -t, --token string        access_token
  ```
