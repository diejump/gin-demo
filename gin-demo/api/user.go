package api

import (
	"bufio"
	"gin-demo/api/middleware"
	"gin-demo/dao"
	"gin-demo/model"
	"gin-demo/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func register(c *gin.Context) {

	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespSuccess(c, "verification failed")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	date := c.PostForm("date") //为了找回密码需要注册时输入一个日期

	flag := dao.SelectUser(username)
	if flag {
		utils.RespFail(c, "user already exists")
		return
	}

	dao.AddUser(username, password, date) //添加用户

	utils.RespSuccess(c, "add user successful")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	dao.AllUser() //更新数据

	username := c.PostForm("username")
	password := c.PostForm("password")

	flag := dao.SelectUser(username)
	if !flag {
		utils.RespFail(c, "user doesn't exists")
		return
	}

	selectPassword := dao.SelectPasswordFromUsername(username)

	if selectPassword != password {
		print(selectPassword)
		utils.RespFail(c, "wrong password")
		return
	}

	claim := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "xxx",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, tokenString)
}

func FindPassword(c *gin.Context) { //找回密码
	if err := c.ShouldBind(&model.FindPassword{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}

	dao.AllUser()
	username := c.PostForm("username")
	date := c.PostForm("date")
	password := dao.GetPassword(username, date)

	if password == "" {
		utils.RespFail(c, "Error")
		return
	}

	utils.RespSuccess(c, password)
}

func ChangePassword(c *gin.Context) {
	if err := c.ShouldBind(&model.ChangePassword{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}

	dao.AllUser()
	date := c.PostForm("date")
	currentPassword := c.PostForm("currentpassword")
	newPassword := c.PostForm("newpassword")

	flag := dao.ChangePassword(date, currentPassword, newPassword)

	if !flag {
		utils.RespFail(c, "change password failed")
	} else {
		utils.RespSuccess(c, "change password successful")
	}

}

func Comments(c *gin.Context) { //简易留言板
	if err := c.ShouldBind(&model.Comments{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}

	comments := c.PostForm("comments")
	f, _ := os.OpenFile("commments.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	writer := bufio.NewWriter(f)
	writer.WriteString(comments)
	writer.WriteString("\n")
	writer.Flush()
	utils.RespSuccess(c, "Thanks for commenting")
}

func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}
