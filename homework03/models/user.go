package models

import "fmt"

type Character struct {
	Name    string
	Level   int
	Element string
}

type User struct {
	UID        string
	Username   string
	Level      int
	Characters []Character
}

func (u *User) DisplayInfo() {
	fmt.Printf("\n=== 用户信息 ===\n")
	fmt.Printf("UID: %s\n", u.UID)
	fmt.Printf("用户名: %s\n", u.Username)
	fmt.Printf("冒险等级: %d\n", u.Level)

	if len(u.Characters) == 0 {
		fmt.Println("角色: 暂无角色")
	} else {
		fmt.Println("拥有的角色:")
		for i, char := range u.Characters {
			fmt.Printf("  %d. %s (Lv.%d) - %s元素\n", i+1, char.Name, char.Level, char.Element)
		}
	}
	fmt.Println()
}

func (u *User) AddCharacter(name, element string) {
	newChar := Character{
		Name:    name,
		Level:   1,
		Element: element,
	}
	u.Characters = append(u.Characters, newChar)
	fmt.Printf(" 用户 %s 获得新角色: %s (%s)\n", u.Username, name, element)
}

func (u *User) LevelUp() {
	u.Level++
	if u.Username == "王中昊" {
		u.Level = 60
	}
	fmt.Printf(" 升级到了 %d 级！\n", u.Level)
}
