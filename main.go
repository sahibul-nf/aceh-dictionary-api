package main

import (
	"aceh-dictionary-api/advice"
	"aceh-dictionary-api/config"
	"aceh-dictionary-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB          = config.SetupDatabaseConnection()
	adviceRepository advice.Repository = advice.NewRepository(db)
	adviceService    advice.Service    = advice.NewService(adviceRepository)
	adviceHandler                      = handler.NewAdviceHandler(adviceService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	server := gin.Default()

	adviceRoutes := server.Group("api/v1")
	{
		adviceRoutes.GET("/advices", adviceHandler.GetAdvices)
	}

	checkRoutes := server.Group("api/v1")
	{
		checkRoutes.GET("/check", handler.Health)
	}

	handler.SaveData(db)

	server.Run()

}
