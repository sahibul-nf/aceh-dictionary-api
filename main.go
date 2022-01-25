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

	adviceRoutes := server.Group("/api/v1")
	{
		adviceRoutes.GET("/advices", adviceHandler.GetAdvices)
	}

	server.Run()

	// SCRAPING WORD BOOK (DATA) AND INSERT TO DB
	// dictRepository := dictionary.NewRepository(db)
	// dictService := dictionary.NewService(dictRepository)

	// fmt.Println("Start scraping data \n -------")
	// data := scraping.FetchAcehIndoDictionary()
	// fmt.Println("Done scraping data")

	// isSuccess, err := dictService.SaveData(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(isSuccess)
	// if isSuccess {
	// 	fmt.Println("Successfully scraping data")
	// }
}
