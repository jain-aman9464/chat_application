package service

import (
	"fmt"
	"github.com/tokopedia/test/chat_app/src/model"
	"github.com/tokopedia/test/chat_app/src/repo"
)

type GroupService struct {
	groupRepo repo.GroupRepo
}

func NewGroupService(groupRep repo.GroupRepo) GroupService {
	return GroupService{groupRepo: groupRep}
}

func (g GroupService) CreateGroup(group model.Group, admin model.User) {
	if !admin.IsAdmin() {
		fmt.Println("You dont have necessary access right to create the group")
		return
	}

	g.groupRepo.CreateGroup(group, admin)
}

func (g GroupService) JoinGroup(group model.Group, user model.User) {
	g.groupRepo.JoinGroup(group, user)
}

func (g GroupService) LeaveGroup(group model.Group, user model.User) {
	g.groupRepo.LeaveGroup(group, user)
}
