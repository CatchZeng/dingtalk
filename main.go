package main

import (
	"log"
	"os"

	"github.com/CatchZeng/dingtalk/cmd/dingtalk"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	dingtalk.Execute()
}
