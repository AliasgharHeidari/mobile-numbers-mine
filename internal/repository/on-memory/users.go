package onmemory

import (
	"log"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/model"
	dataonredis "github.com/AliasgharHeidari/mobile-numbers-mine/internal/repository/redis"
)

func InitUsers() {

	initUsers := []model.User{
		{
			ID:         1,
			Name:       "Ali",
			FamilyName: "Heidari",
			Age:        18,
			IsMarried:  false,
		},
		{
			ID:         2,
			Name:       "Amir",
			FamilyName: "Barkhordari",
			Age:        21,
			IsMarried:  false,
		},
	}

	for _, user := range initUsers {
		err := dataonredis.SaveUserToRedis(user)
		if err != nil {
			log.Printf("failed to save user %s to redis, error: %+v", user.Name, err)
		}
	}
}
