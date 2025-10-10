package onmemory

import (
	"sync"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/model"
)

type UserData struct {
	Users           []model.User
	UserChangeMutex *sync.Mutex
}

var UsersRepo UserData

func LoadInitUsers() {
	UsersRepo = UserData{}
	UsersRepo.Users = []model.User{
		{
			ID:        1,
			Name:      "Alice",
			IsMarried: false,
		},
		{
			ID:        2,
			Name:      "Bob",
			IsMarried: true,
		},
	}

	UsersRepo.UserChangeMutex = &sync.Mutex{}

}
