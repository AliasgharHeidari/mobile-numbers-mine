package onmemory

import "github.com/AliasgharHeidari/mobile-numbers-mine/internal/model"

var Users []model.User

func LoadInitUsers() {
	Users = append(Users, model.User{
		ID:         1,
		Name:       "Ashva",
		FamilyName: "Patel",
		Age:        24,
		IsMarried:  false})
}
