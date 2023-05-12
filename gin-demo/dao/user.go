package dao

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type User struct {
	Username string
	Password string
	date     string
}

var user [100]User

func AddUser(username, password string, date string) {

	for i := 0; i < 100; i++ {
		if user[i].Username == "" {
			user[i].Username = username
			user[i].Password = password
			user[i].date = date
			break
		}
	}

	f := "user.txt"
	f1, _ := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	writer := bufio.NewWriter(f1)
	writer.WriteString(username)
	writer.WriteString("\n") //换行符会被读取到
	writer.WriteString(password)
	writer.WriteString("\n")
	writer.WriteString(date)
	writer.WriteString("\n")
	writer.Flush()
	f1.Close()
}

func AllUser() { //根据user.txt同步所有数据
	f, _ := os.OpenFile("user.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	reader := bufio.NewReader(f)
	i := 0
	for {
		username, _ := reader.ReadString('\n')
		password, _ := reader.ReadString('\n')
		date, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		username = strings.TrimSpace(username) //去除username中的换行符
		password = strings.TrimSpace(password)
		date = strings.TrimSpace(date)

		user[i].Username = username
		user[i].Password = password
		user[i].date = date
		i++
	}

	for j := 0; user[j].Username != ""; j++ {
		print(user[j].Username, user[j].Password, user[j].date)
	}
}

func SelectUser(username string) bool {
	for i := 0; i < 100; i++ {
		if strings.Compare(user[i].Username, username) == 0 {
			return true
		}
	}
	return false
}

func SelectPasswordFromUsername(username string) string {
	for i := 0; i < 100; i++ {
		if strings.Compare(user[i].Username, username) == 0 {
			return user[i].Password
		}
	}
	return ""
}

func GetPassword(username, date string) string { //通过用户名和日期得到密码
	for i := 0; i < 100; i++ {
		if strings.Compare(user[i].Username, username) == 0 {
			if strings.Compare(user[i].date, date) == 0 {
				return user[i].Password
			}
		}
	}
	return ""
}

func ChangePassword(date, currentPassword, newPassword string) bool { //修改密码
	for i := 0; i < 100; i++ {
		if strings.Compare(user[i].Password, currentPassword) == 0 {
			if strings.Compare(user[i].date, date) == 0 {
				if strings.Compare(newPassword, currentPassword) != 0 {
					user[i].Password = newPassword
					f, _ := os.OpenFile("user.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
					writer := bufio.NewWriter(f)
					for j := 0; user[j].Username != ""; j++ {
						//实在没找到稳定替换文件里某一行数据的方法，只能覆盖全写一遍
						writer.WriteString(user[j].Username)
						writer.WriteString("\n")
						writer.WriteString(user[j].Password)
						writer.WriteString("\n")
						writer.WriteString(user[j].date)
						writer.WriteString("\n")
						print(user[j].Username, user[j].Password, user[j].date)
						writer.Flush()
					}
					return true
				}
			}
		}
	}
	return false
}
