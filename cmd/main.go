package main

import (
	apiserver "github.com/AliasgharHeidari/mobile-numbers-mine/internal/api/server"
	onmemory "github.com/AliasgharHeidari/mobile-numbers-mine/internal/repository/on-memory"
)

func main() {
	onmemory.LoadInitUsers()
	apiserver.Start()
}
