package service

import (
	"errors"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/model"
	onmemory "github.com/Golang-Training-entry-3/mobile-numbers/internal/repository/on-memory"
)

func GetUserList() ([]model.User, error) {
	users := onmemory.Users
	return users, nil
}

func GetUserByID(id int) (model.User, error) {
	for _, user := range onmemory.Users {
		if user.ID == id {
			return user, nil
		}
	}

	errorMessage := errors.New("User not found")
	return model.User{}, errorMessage

}

func CreateUser(user model.User) (int, error) {
	newRandomId := len(onmemory.Users) + 1

	user.ID = newRandomId

	onmemory.Users = append(onmemory.Users, user)
	return newRandomId, nil
}
