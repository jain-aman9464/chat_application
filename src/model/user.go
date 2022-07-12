package model

import "github.com/google/uuid"

type User struct {
	userID  string
	name    string
	email   string
	isAdmin bool
}

func NewUser(name, email string, isAdmin bool) User {
	return User{
		userID: uuid.New().String(),
		name:   name,
		email:  email,
	}
}

func (u User) SetName(name string) {
	u.name = name
}

func (u User) SetEmail(email string) {
	u.email = email
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetUserID() string {
	return u.userID
}

func (u User) GetEmail() string {
	return u.email
}

func (u User) IsAdmin() bool {
	return u.isAdmin
}

func (u User) SetAdmin() {
	u.isAdmin = true
}
