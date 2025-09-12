package main

import (
	apiserver "github.com/Golang-Training-entry-3/mobile-numbers/internal/api/server"
	onmemory "github.com/Golang-Training-entry-3/mobile-numbers/internal/repository/on-memory"
)

func main() {
	onmemory.LoadInitUsers()
	apiserver.Start()
}
