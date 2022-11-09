package main

import (
	"aceh-dictionary-api/config"
	"aceh-dictionary-api/dictionary"
	"aceh-dictionary-api/handler"
	"aceh-dictionary-api/search"

	"github.com/gin-gonic/gin"
)

var (
	db               = config.SetupDatabaseConnection()
	searchRepository = search.NewRepository(db)
	searchService    = search.NewService(searchRepository)
	searchHandler    = handler.NewSearchHandler(searchService)

	dictionaryRepository = dictionary.NewRepository(db)
	dictionaryService    = dictionary.NewService(dictionaryRepository)
	dictionaryHandler    = handler.NewDictionaryHandler(dictionaryService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(CORSMiddleware())

	searchRoutes := server.Group("api/v1")
	{
		searchRoutes.GET("/search", searchHandler.Search)
	}

	dictionaryRoutes := server.Group("api/v1")
	{
		dictionaryRoutes.POST("/dictionaries", dictionaryHandler.AddNewDictionary)
		dictionaryRoutes.GET("/dictionaries", dictionaryHandler.GetDictionaries)
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
