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

func (us *UserService) InitExampleData() {

	us.Register("1001", "旅行者")
	us.Register("1002", "派蒙")
	us.Register("1003", "王中昊")

	us.AddCharacter("1001", "荧", "无")
	us.AddCharacter("1001", "凯亚", "冰")
	us.AddCharacter("1002", "派蒙", "无")
	us.AddCharacter("1003", "甘雨", "冰")
	us.AddCharacter("1003", "胡桃", "火")
	us.AddCharacter("1003", "芙宁娜", "水")
	us.AddCharacter("1003", "流萤", "火")
	us.AddCharacter("1003", "雷电将军", "雷")

	us.LevelUp("1001")
	us.LevelUp("1001")
	us.LevelUp("1002")
	us.LevelUp("1003")

	fmt.Println("示例数据已加载完成！")
}
