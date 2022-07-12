package repo

import (
	"fmt"
	"github.com/tokopedia/test/chat_app/src/model"
)

type GroupRepo struct {
	groupMap map[string][]model.User
	adminMap map[string]int
}

func NewGroupRepo() GroupRepo {
	return GroupRepo{
		groupMap: make(map[string][]model.User, 100),
		adminMap: make(map[string]int, 100),
	}
}

func (g GroupRepo) CreateGroup(group model.Group, admin model.User) {
	if _, ok := g.groupMap[group.GetGroupID()]; ok {
		fmt.Printf("%s Group already exists\n", group.GetGroupName())
		return
	}

	if count, ok := g.adminMap[admin.GetUserID()]; ok && count > 3 {
		fmt.Printf("%s Admin not eligible to create another group. Max limit exceeded\n", admin.GetName())
		return
	}

	if _, ok := g.adminMap[admin.GetUserID()]; !ok {
		g.adminMap[admin.GetUserID()] = 0
	}

	arr := make([]model.User, 100)
	arr = append(arr, admin)

	g.groupMap[group.GetGroupID()] = arr
	g.adminMap[admin.GetUserID()] += g.adminMap[admin.GetUserID()] + 1

	fmt.Printf("%s Group created\n", group.GetGroupName())
}

func (g GroupRepo) JoinGroup(group model.Group, user model.User) {
	if len(group.GetUsers()) >= group.GetCapacity() {
		fmt.Printf("%s group already full\n", group.GetGroupName())
		return
	}

	usersInGroup := g.groupMap[group.GetGroupID()]
	for _, u := range usersInGroup {
		if u.GetUserID() == user.GetUserID() {
			fmt.Printf("%s user already exists in the group\n", user.GetName())
			return
		}
	}

	usersInGroup = append(usersInGroup, user)
	g.groupMap[group.GetGroupID()] = usersInGroup
	group.SetUsers(usersInGroup)
	fmt.Printf("%s user joined the group %s \n", user.GetName(), group.GetGroupName())
}

func (g GroupRepo) LeaveGroup(group model.Group, user model.User) {
	if len(group.GetUsers()) == 0 {
		fmt.Printf("%s group already empty \n", group.GetGroupName())
		return
	}

	usersInGroup := g.groupMap[group.GetGroupID()]
	// Transfer admin rights
	if user.IsAdmin() {
		usersInGroup[1].SetAdmin()
	}

	for idx, u := range usersInGroup {
		if u.GetUserID() == user.GetUserID() {
			usersInGroup = append(usersInGroup[:idx], usersInGroup[idx+1])
			break
		}
	}

	g.groupMap[group.GetGroupID()] = usersInGroup

	fmt.Printf("%s user left the group %s \n", user.GetName(), group.GetGroupName())
}
