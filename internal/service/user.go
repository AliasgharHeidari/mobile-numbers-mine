package service

import (
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/model"
	onmemory "github.com/Golang-Training-entry-3/mobile-numbers/internal/repository/on-memory"
)

func GetUserList() ([]model.User, error) {
	users := onmemory.Users
	return users, nil
}
