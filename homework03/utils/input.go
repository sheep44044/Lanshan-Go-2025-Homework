package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func ShowMenu() {
	fmt.Println("\n=== 原神用户管理系统 ===")
	fmt.Println("1. 注册用户")
	fmt.Println("2. 添加角色")
	fmt.Println("3. 用户升级")
	fmt.Println("4. 查看用户信息")
	fmt.Println("5. 查看所有用户")
	fmt.Println("6. 退出")
	fmt.Print("请选择操作: ")
}
