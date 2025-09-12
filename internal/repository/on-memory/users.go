package onmemory

import "github.com/Golang-Training-entry-3/mobile-numbers/internal/model"

var Users []model.User

func LoadInitUsers() {
	Users = append(Users, model.User{
		ID:         1,
		Name:       "Ashva",
		FamilyName: "Patel",
		Age:        24,
		IsMarried:  false})
}
