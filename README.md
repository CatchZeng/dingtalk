# dingtalk

[English](https://github.com/CatchZeng/dingtalk/blob/master/READMEEN.md)

> DingTalk(dingding) 是钉钉机器人的 go 实现。支持 `Docker`、`命令行`模式，`module` 模式，支持`加签`安全设置，支持`链式语法`创建消息，支持`文本、链接、Markdown` 消息类型

## 文档

[钉钉文档](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)

## 特性

- [x] 支持[Docker](https://github.com/CatchZeng/dingtalk#Docker)

- [x] 支持[module](https://github.com/CatchZeng/dingtalk#%E4%BD%9C%E4%B8%BA-module)

- [x] 支持[命令行模式](https://github.com/CatchZeng/dingtalk#%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%B7%A5%E5%85%B7)

- [x] 支持加签

  ![sign](https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1572261283991-f8e35f4d-6997-4a02-9704-843ee8f97464.png)

- [x] Text 消息

  ![text](https://img.alicdn.com/tfs/TB1jFpqaRxRMKJjy0FdXXaifFXa-497-133.png)

- [x] Link 消息

  ![link](https://dingtalkdoc.oss-cn-beijing.aliyuncs.com/images/0.0.210/1570679827267-6243216b-d1c3-48b7-9b1e-0f0b4211b50b.png)

- [x] Markdown 消息

  ![markdown](https://img.alicdn.com/tfs/TB1yL3taUgQMeJjy0FeXXXOEVXa-492-380.png)

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

### Docker

```shell
docker run catchzeng/dingtalk dingtalk text -t 1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f -s SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68 -c "docker test"
```

### 作为 module

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

### 命令行工具

#### Demo

```shell
dingtalk text -t 1c53e149ba5de6597cxxxxxx0e901fdxxxxxx80b8ac141e4a75afdc44c85ca4f -s SECb90923e19e58b466481e9e7b7a5bxxxxxx4531axxxxxxad3967fb29f0eae5c68 -c "测试命令行 & at 某个人" -m ["177010xxx60"]
```

#### Help

- dingtalk

  ```shell
  $ dingtalk -h
  dingtalk is a command line tool for DingTalk

  Usage:
    dingtalk [command]

  Available Commands:
    help        Help about any command
    link        send link message with DingTalk robot
    markdown    send markdown message with DingTalk robot
    text        send text message with DingTalk robot

  Flags:
    -m, --atMobiles stringArray   atMobiles
    -h, --help                    help for dingtalk
    -a, --isAtAll                 isAtAll
    -s, --secret string           secret
    -t, --token string            access_token

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
    -m, --atMobiles stringArray   atMobiles
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
    -m, --atMobiles stringArray   atMobiles
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
    -m, --atMobiles stringArray   atMobiles
    -a, --isAtAll                 isAtAll
    -s, --secret string           secret
    -t, --token string            access_token
  ```
