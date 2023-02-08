package auth

import (
	"goep1/orm"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Binding from JSON
type RegisterBody struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Check User
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Exists"})
		return
	}

	//สร้าง User
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := orm.User{Username: json.Username,Password: string(encryptedPassword),Fullname: json.Fullname,Avatar:   json.Avatar}
	orm.Db.Create(&user) // pass pointer of data to Create
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "userId": user.ID, "message": "User Create Success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "userId": user.ID, "message": "User Create Fail"})
	}
}
