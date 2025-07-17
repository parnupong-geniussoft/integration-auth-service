package usecases

import (
	"fmt"
	"go-fiber-clean-arch-example/modules/entities"
	"go-fiber-clean-arch-example/modules/users/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UsersUsecase interface {
	Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error)
}

type usersUsecase struct {
	UsersRepo repositories.UsersRepository
}

// Constructor
func NewUsersUsecase(usersRepo repositories.UsersRepository) UsersUsecase {
	return &usersUsecase{
		UsersRepo: usersRepo,
	}
}

func (u *usersUsecase) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	// Hash a password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	// Send req next to repository
	user, err := u.UsersRepo.Register(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
