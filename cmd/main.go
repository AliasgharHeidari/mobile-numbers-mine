// @title Mobile Numbers Mine API
// @version 1.0
// @description API for managing users and their mobile numbers.
// @contact.name Aliasghar
// @host localhost:9898
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	apiserver "github.com/AliasgharHeidari/mobile-numbers-mine/internal/api/server"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/config"
	onmemory "github.com/AliasgharHeidari/mobile-numbers-mine/internal/repository/on-memory"
)

func main() {
	config.InitConfig()
	onmemory.LoadInitUsers()
	apiserver.Start()
}