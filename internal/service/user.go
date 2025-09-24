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

	errorMessage := errors.New("user not found")
	return model.User{}, errorMessage

}

func CreateUser(user model.User) (int, error) {
	newRandomId := len(onmemory.Users) + 1

	user.ID = newRandomId

	onmemory.Users = append(onmemory.Users, user)
	return newRandomId, nil
}
func UpdateUserByID(id int, updatedUser model.User) error {
	for i, user := range onmemory.Users {
		if user.ID == id {
			updatedUser.ID = id 
			onmemory.Users[i] = updatedUser
			return nil
		}
	}
	return errors.New("user not found")
}

func DeleteUserByID (id int) error {
	for i, user := range onmemory.Users{
		if user.ID == id {
			onmemory.Users = append(onmemory.Users[:i], onmemory.Users[i+1:]...)
			return nil
		}
	}

	return  errors.New("user not found")

}