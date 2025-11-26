package dao

import (
	"os"
	"strings"
)

var database = make(map[string]string)

const dataFile = "users.txt"

// 初始化时从文件加载数据
func init() {
	loadUsersFromFile()
}

func loadUsersFromFile() {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在，创建空文件
			saveUsersToFile()
			return
		}
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			database[parts[0]] = parts[1]
		}
	}
}

func saveUsersToFile() {
	var lines []string
	for username, password := range database {
		lines = append(lines, username+":"+password)
	}

	data := strings.Join(lines, "\n")
	os.WriteFile(dataFile, []byte(data), 0644)
}

func AddUser(username string, password string) {
	database[username] = password
	saveUsersToFile()
}

func FindUser(username string, password string) bool {
	if pwd, ok := database[username]; ok {
		return pwd == password
	}
	return false
}

func ModifyPassword(username string, password string) {
	database[username] = password
	saveUsersToFile()
}
