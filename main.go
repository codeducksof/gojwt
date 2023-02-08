package main

import (
	AuthController "goep1/controller/auth"
	"goep1/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	orm.InitDB()
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.Run()
}
