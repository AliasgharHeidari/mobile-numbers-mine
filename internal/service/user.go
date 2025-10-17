package service

import (
	"errors"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/model"
	onmemory "github.com/AliasgharHeidari/mobile-numbers-mine/internal/repository/on-memory"
)

func GetUserList() ([]model.User, error) {
	users := onmemory.UsersRepo.Users
	return users, nil
}

func GetUserByID(id int) (model.User, error) {
	for _, user := range onmemory.UsersRepo.Users {
		if user.ID == id {
			return user, nil
		}
	}

	errorMessage := errors.New("user not found")
	return model.User{}, errorMessage

}

func CreateUser(user model.User) (int, error) {
	onmemory.UsersRepo.UserChangeMutex.Lock()
	defer onmemory.UsersRepo.UserChangeMutex.Unlock()
	newRandomId := len(onmemory.UsersRepo.Users) + 1

	user.ID = newRandomId

	onmemory.UsersRepo.Users = append(onmemory.UsersRepo.Users, user)
	return newRandomId, nil
}
func UpdateUserByID(id int, updatedUser model.User) error {
	onmemory.UsersRepo.UserChangeMutex.Lock()
	defer onmemory.UsersRepo.UserChangeMutex.Unlock()
	for i, user := range onmemory.UsersRepo.Users {
		if user.ID == id {
			updatedUser.ID = id
			onmemory.UsersRepo.Users[i] = updatedUser
			return nil
		}
	}
	return errors.New("user not found")
}

func DeleteUserByID(id int) error {
	for i, user := range onmemory.UsersRepo.Users {
		if user.ID == id {
			onmemory.UsersRepo.Users = append(onmemory.UsersRepo.Users[:i], onmemory.UsersRepo.Users[i+1:]...)
			return nil
		}
	}

	return errors.New("user not found")

}

func AddMobileNumber(id int, mobileNumbers model.MobileNumber) error {
	for i, user := range onmemory.UsersRepo.Users {
		if user.ID == id {
			onmemory.UsersRepo.Users[i].MobileNumbers = append(onmemory.UsersRepo.Users[i].MobileNumbers, mobileNumbers)
			return nil
		}
	}

	return errors.New("user not found")

}

func DeleteMobileNumber(id int, Number string) error {
	for i, user := range onmemory.UsersRepo.Users {
		if user.ID == id {
			for j, mobile := range user.MobileNumbers {
				if mobile.Number == Number {
					onmemory.UsersRepo.Users[i].MobileNumbers = append(user.MobileNumbers[:j], user.MobileNumbers[j+1:]...)
					return nil
				}
			}
			return errors.New("MobileNumber not found")
		}
	}
	return errors.New("error : user not found")
}
