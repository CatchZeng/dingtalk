package main

import (
	"github.com/CatchZeng/dingtalk/cmd/dingtalk"
	"log"
)

func main() {
	log.SetFlags(0)
	dingtalk.Execute()
}
