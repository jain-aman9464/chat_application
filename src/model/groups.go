package model

import (
	"github.com/google/uuid"
)

type Group struct {
	groupID   string
	groupName string
	capacity  int
	users     []User
}

func NewGroup(name string, capacity int, admin User) Group {
	return Group{
		groupID:   uuid.New().String(),
		groupName: name,
		capacity:  capacity,
		users:     make([]User, 0),
		//admin:     admin,
	}
}

func (g Group) GetGroupID() string {
	return g.groupID
}

func (g Group) GetCapacity() int {
	return g.capacity
}

func (g Group) GetGroupName() string {
	return g.groupName
}

func (g Group) GetUsers() []User {
	return g.users
}

func (g *Group) SetUsers(users []User) {
	g.users = users
}
