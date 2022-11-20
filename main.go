package main

import (
	"aceh-dictionary-api/auth"
	"aceh-dictionary-api/bookmark"
	"aceh-dictionary-api/config"
	"aceh-dictionary-api/dictionary"
	"aceh-dictionary-api/handler"
	"aceh-dictionary-api/handler/middleware"
	"aceh-dictionary-api/search"
	"aceh-dictionary-api/unsplash"
	"aceh-dictionary-api/user"

	"github.com/gin-gonic/gin"
)

var (
	db = config.SetupDatabaseConnection()

	unsplashRepository = unsplash.NewRepository()

	searchRepository = search.NewRepository(db)
	searchService    = search.NewService(searchRepository)
	searchHandler    = handler.NewSearchHandler(searchService)

	dictionaryRepository = dictionary.NewRepository(db)
	dictionaryService    = dictionary.NewService(dictionaryRepository, unsplashRepository)
	dictionaryHandler    = handler.NewDictionaryHandler(dictionaryService)

	userRepository = user.NewRepository(db)
	userService    = user.NewService(userRepository)
	authService    = auth.NewService()
	userHandler    = handler.NewUserHandler(userService, authService)

	bookmarkRepository = bookmark.NewRepository(db)
	bookmarkService    = bookmark.NewService(bookmarkRepository)
	bookmarkHandler    = handler.NewBookmarkHandler(bookmarkService)
)

func main() {
	// ! Comment this line if you wanna deploy to Heroku
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	panic(err)
	// }

	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	searchRoutes := server.Group("api/v1")
	{
		searchRoutes.GET("/search", searchHandler.Search)
	}

	dictionaryRoutes := server.Group("api/v1")
	{
		dictionaryRoutes.POST("/dictionaries", dictionaryHandler.AddNewWord)
		dictionaryRoutes.GET("/dictionaries", dictionaryHandler.GetWords)
		dictionaryRoutes.GET("/dictionaries/:id", dictionaryHandler.GetWord)
	}

	userRoutes := server.Group("api/v1")
	{
		userRoutes.POST("/users", userHandler.RegisterUser)
		userRoutes.POST("/users/sessions", userHandler.Login)
	}

	bookmarkRoutes := server.Group("api/v1")
	{
		bookmarkRoutes.POST("/bookmarks", middleware.AuthMiddleware(authService, userService), bookmarkHandler.MarkedAndUnmarkedWord)
		bookmarkRoutes.GET("/bookmarks", middleware.AuthMiddleware(authService, userService), bookmarkHandler.GetMarkedWordsByUserID)
		bookmarkRoutes.GET("/bookmark", middleware.AuthMiddleware(authService, userService), bookmarkHandler.GetMarkedWords)
		bookmarkRoutes.DELETE("/bookmarks/:id", middleware.AuthMiddleware(authService, userService), bookmarkHandler.DeleteMarkedWord)
		bookmarkRoutes.DELETE("/bookmarks", middleware.AuthMiddleware(authService, userService), bookmarkHandler.DeleteAllMarkedWordsByUserID)
	}

	server.GET("/", handler.Index)
	server.Run()
}
