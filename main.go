package main

import (
	"aceh-dictionary-api/advice"
	"aceh-dictionary-api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/kamus_aceh?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	adviceRepository := advice.NewRepository(db)
	adviceService := advice.NewService(adviceRepository)
	adviceHandler := handler.NewAdviceHandler(adviceService)

	api := gin.Default()
	v1 := api.Group("/api/v1/")

	v1.GET("/advices", adviceHandler.GetAdvices)

	api.Run()

	// SCRAPING WORD BOOK (DATA) AND INSERT TO DB
	// dictRepository := dictionary.NewRepository(db)
	// dictService := dictionary.NewService(dictRepository)

	// data := scraping.FetchAcehIndoDictionary()

	// isSuccess, err := dictService.SaveData(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(isSuccess)
}
