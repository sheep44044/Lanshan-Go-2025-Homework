package initializers

import (
	"awesomeProject1/homework03/services"
	"fmt"
)

func InitExampleData(us *services.UserService) {

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
