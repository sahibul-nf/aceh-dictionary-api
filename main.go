package main

import (
	"aceh-dictionary-api/advice"
	"aceh-dictionary-api/config"
	"aceh-dictionary-api/dictionary"
	"aceh-dictionary-api/handler"

	"github.com/gin-gonic/gin"
)

var (
	db               = config.SetupDatabaseConnection()
	adviceRepository = advice.NewRepository(db)
	adviceService    = advice.NewService(adviceRepository)
	adviceHandler    = handler.NewAdviceHandler(adviceService)

	dictRepository = dictionary.NewRepository(db)
	dictService    = dictionary.NewService(dictRepository)
	dictHandler    = handler.NewDictionaryHandler(dictService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	server := gin.Default()

	adviceRoutes := server.Group("api/v1")
	{
		adviceRoutes.POST("/advices", adviceHandler.GetAdvices)
	}

	dictionaryRoutes := server.Group("api/v1")
	{
		dictionaryRoutes.POST("/dictionary", dictHandler.CreateDictionaryData)
	}

	checkRoutes := server.Group("api/v1")
	{
		checkRoutes.GET("/check", handler.Health)
	}

	server.Run()
}
