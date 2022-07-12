package service

import (
	"github.com/tokopedia/test/chat_app/src/model"
	"github.com/tokopedia/test/chat_app/src/repo"
)

type UserService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return UserService{userRepo: userRepo}
}

func (u UserService) Signup(user model.User) {
	u.userRepo.Signup(user)
}
