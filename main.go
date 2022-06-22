package main

import (
	"aceh-dictionary-api/advice"
	"aceh-dictionary-api/config"
	"aceh-dictionary-api/handler"
	"aceh-dictionary-api/vocabulary"

	"github.com/gin-gonic/gin"
)

var (
	db               = config.SetupDatabaseConnection()
	adviceRepository = advice.NewRepository(db)
	adviceService    = advice.NewService(adviceRepository)
	adviceHandler    = handler.NewAdviceHandler(adviceService)

	vocabRepository = vocabulary.NewRepository(db)
	vocabService    = vocabulary.NewService(vocabRepository)
	vocabHandler    = handler.NewVocabularyHandler(vocabService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(CORSMiddleware())

	adviceRoutes := server.Group("api/v1")
	{
		adviceRoutes.GET("/advices", adviceHandler.GetAdvices)
	}

	dictionaryRoutes := server.Group("api/v1")
	{
		dictionaryRoutes.POST("/vocabularies", vocabHandler.AddNewVocabulary)
		dictionaryRoutes.GET("/vocabularies", vocabHandler.GetAllVocabularyData)
	}

	server.GET("/", handler.Index)
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
