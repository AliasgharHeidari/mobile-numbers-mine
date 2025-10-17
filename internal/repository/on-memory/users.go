package onmemory

import (
	"sync"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/model"
)
type UserData struct {
	Users []model.User
	UserChangeMutex *sync.Mutex
}

var UsersRepo UserData

func LoadInitUsers() {
	UsersRepo = UserData{}
	UsersRepo.Users = []model.User{
		{
			ID:     1,
			Name:   "Ali",
			FamilyName: "Heidari",
			Age: 18,
			IsMarried: false,
		},

		{
			ID :   	2,
			Name: "Amir",
			FamilyName: "Barkhordari",
			Age: 21,
			IsMarried: false,
		},
	}

	UsersRepo.UserChangeMutex = &sync.Mutex{}

}
