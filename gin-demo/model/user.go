package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Date     string `form:"date" json:"date"`
}

type FindPassword struct {
	Username string `form:"username" json:"username" binding:"required"`
	Date     string `form:"date" json:"date" binding:"required"`
}

type ChangePassword struct {
	Date            string `form:"date" json:"date" binding:"required"`
	CurrentPassword string `form:"currentpassword" json:"currentpassword" binding:"required"`
	NewPassword     string `form:"newpassword" json:"newpassword" binding:"required"`
}

type Comments struct {
	Comments string `form:"comments" json:"comments" binding:"required"`
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
