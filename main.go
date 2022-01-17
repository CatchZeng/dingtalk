package main

import (
	"log"

	"github.com/CatchZeng/dingtalk/cmd/dingtalk"
)

func main() {
	log.SetFlags(0)
	dingtalk.Execute()
}
