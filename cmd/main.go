// @title Mine API - Swagger Docs
// @version 1.0
// @description API for managing users and their mobile numbers.
// @contact.name Aliasghar Heidari
// @contact.email ali.heidari@gmail.com
// @host 127.0.0.1:9898
// @BasePath /
// @schemes http

// Security Definitions
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// Apply security globally
// @Security BearerAuth

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