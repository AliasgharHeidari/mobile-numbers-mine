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
