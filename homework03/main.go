package main

import (
	"awesomeProject1/homework03/initializers"
	"awesomeProject1/homework03/services"
	"awesomeProject1/homework03/utils"
	"fmt"
)

func main() {

	userService := services.NewUserService()

	initializers.InitExampleData(userService)

	fmt.Println("🎮 欢迎使用原神用户管理系统!")

	for {
		utils.ShowMenu()
		choice := utils.GetInput("")

		switch choice {
		case "1":
			handleRegister(userService)
		case "2":
			handleAddCharacter(userService)
		case "3":
			handleLevelUp(userService)
		case "4":
			handleShowUser(userService)
		case "5":
			userService.ShowAllUsers()
		case "6":
			fmt.Println("感谢使用，再见！")
			return
		default:
			fmt.Println(" 无效选择，请重新输入")
		}
	}
}

func handleRegister(service *services.UserService) {
	uid := utils.GetInput("输入用户ID: ")
	username := utils.GetInput("输入用户名: ")

	err := service.Register(uid, username)
	if err != nil {
		fmt.Printf(" 注册失败: %v\n", err)
	}
}

func handleAddCharacter(service *services.UserService) {
	uid := utils.GetInput("输入用户ID: ")
	charName := utils.GetInput("输入角色名: ")
	element := utils.GetInput("输入元素属性: ")

	err := service.AddCharacter(uid, charName, element)
	if err != nil {
		fmt.Printf(" 添加角色失败: %v\n", err)
	}
}

func handleLevelUp(service *services.UserService) {
	uid := utils.GetInput("输入要升级的用户ID: ")

	err := service.LevelUp(uid)
	if err != nil {
		fmt.Printf(" 升级失败: %v\n", err)
	}
}

func handleShowUser(service *services.UserService) {
	uid := utils.GetInput("输入要查看的用户ID: ")

	user, err := service.GetUser(uid)
	if err != nil {
		fmt.Printf(" 查看用户失败: %v\n", err)
		return
	}

	user.DisplayInfo()
}
