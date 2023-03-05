package main

import (
	"soul/pkg/config"
	"soul/pkg/logger"
	"soul/pkg/server"
)

func init() {
	config.LoadConfig()
	logger.InitLogger()
}

func main() {
	server.StartServer()
}
