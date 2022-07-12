package repo

import (
	"fmt"
	"github.com/tokopedia/test/chat_app/src/model"
)

type UserRepo struct {
	userMap map[string]model.User
}

func NewUserRepo() UserRepo {
	return UserRepo{
		userMap: make(map[string]model.User, 100),
	}
}

func (u UserRepo) Signup(user model.User) {
	if _, ok := u.userMap[user.GetUserID()]; ok {
		fmt.Printf("User %s already exists in our records \n", user.GetName())
		return
	}

	u.userMap[user.GetUserID()] = user
	fmt.Printf("User %s signed up \n", user.GetName())
}
