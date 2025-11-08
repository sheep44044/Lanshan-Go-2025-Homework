package services

import (
	"awesomeProject1/homework03/models"
	"fmt"
)

type UserService struct {
	users map[string]*models.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*models.User),
	}
}

func (us *UserService) Register(uid, username string) error {
	if _, exists := us.users[uid]; exists {
		return fmt.Errorf("用户ID %s 已存在", uid)
	}

	newUser := &models.User{
		UID:      uid,
		Username: username,
		Level:    1,
	}
	us.users[uid] = newUser
	fmt.Printf(" 用户 %s 注册成功！\n", username)
	return nil
}

func (us *UserService) GetUser(uid string) (*models.User, error) {
	user, exists := us.users[uid]
	if !exists {
		return nil, fmt.Errorf("用户不存在")
	}
	return user, nil
}

func (us *UserService) AddCharacter(uid, charName, element string) error {
	user, exists := us.users[uid]
	if !exists {
		return fmt.Errorf("用户不存在")
	}

	user.AddCharacter(charName, element)
	return nil
}

func (us *UserService) LevelUp(uid string) error {
	user, exists := us.users[uid]
	if !exists {
		return fmt.Errorf("用户不存在")
	}

	user.LevelUp()
	return nil
}

func (us *UserService) ShowAllUsers() {
	if len(us.users) == 0 {
		fmt.Println("暂无用户")
		return
	}

	fmt.Println("\n=== 所有用户 ===")
	for _, user := range us.users {
		fmt.Printf("UID: %s, 用户名: %s, 等级: %d, 角色数: %d\n",
			user.UID, user.Username, user.Level, len(user.Characters))
	}
}
