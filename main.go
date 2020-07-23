package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ramezanpour/users/services"
	"github.com/ramezanpour/users/utilities/database"
)

func main() {
	engine := gin.Default()
	database.Init()
	defer database.Close()
	userService := services.UserService{}
	userService.Serve(engine)

	engine.Run(":3000")
}
