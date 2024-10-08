package main

import (
	_ "github.com/lib/pq"
	"main/config"
	"main/server"
)

func main() {
	accessConfig, refreshConfig := config.TokenConfig()
	server.Start(accessConfig, refreshConfig)
}
