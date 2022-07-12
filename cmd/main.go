package main

import (
	"github.com/tokopedia/test/chat_app/src/model"
	"github.com/tokopedia/test/chat_app/src/repo"
	"github.com/tokopedia/test/chat_app/src/service"
)

func main() {
	userRepo := repo.NewUserRepo()
	groupRepo := repo.NewGroupRepo()

	userService := service.NewUserService(userRepo)
	groupService := service.NewGroupService(groupRepo)

	user1 := model.NewUser("u1", "u1@gmail.com", true)
	user2 := model.NewUser("u2", "u2@gmail.com", false)
	user3 := model.NewUser("u3", "u3@gmail.com", false)
	user4 := model.NewUser("u4", "u4@gmail.com", false)
	user5 := model.NewUser("u5", "u5@gmail.com", false)

	userService.Signup(user1)
	userService.Signup(user1) // return error
	userService.Signup(user2)
	userService.Signup(user3)
	userService.Signup(user4)
	userService.Signup(user5)

	group1 := model.NewGroup("sample_group_1", 10, user1)
	group2 := model.NewGroup("sample_group_2", 10, user1)
	group3 := model.NewGroup("sample_group_3", 10, user1)
	group4 := model.NewGroup("sample_group_4", 10, user1)

	groupService.CreateGroup(group1, user1)
	groupService.CreateGroup(group2, user1)
	groupService.CreateGroup(group3, user1)
	groupService.CreateGroup(group4, user1) // Should return error

	groupService.JoinGroup(group1, user2)
	groupService.JoinGroup(group1, user2) // should return error
	groupService.JoinGroup(group1, user3)
	groupService.JoinGroup(group1, user4)
}
