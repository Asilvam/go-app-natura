package user_service

import (
	"github.com/Asilvam/go-app-natura.git/models"
	user_repository "github.com/Asilvam/go-app-natura.git/repositories/user.repository"
)

func CreateUser(user models.User) error {
	err := user_repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil

}

func GetUsers() (models.Users, error) {

	users, err := user_repository.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}
