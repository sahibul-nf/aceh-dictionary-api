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
	server.Use(CORSMiddleware())

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		/*
		   c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		   c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		   c.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		   c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
		*/

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
