package main

import (
	"github.com/869413421/wechatbot/bootstrap"
	"github.com/869413421/wechatbot/handlers"
)

func main() {
	go handlers.StartOpenApi()
	bootstrap.Run()
}
