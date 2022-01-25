package handler

import (
	"aceh-dictionary-api/dictionary"
	"aceh-dictionary-api/scraping"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Health(ctx *gin.Context) {
	response := map[string]string{
		"message": "ok!",
	}
	ctx.JSON(http.StatusOK, response)
}

func SaveData(db *gorm.DB) {
	// SCRAPING WORD BOOK (DATA) AND INSERT TO DB
	dictRepository := dictionary.NewRepository(db)
	dictService := dictionary.NewService(dictRepository)

	fmt.Println("Start scraping data \n -------")
	data := scraping.FetchAcehIndoDictionary()
	fmt.Println("Done scraping data")

	isSuccess, err := dictService.SaveData(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(isSuccess)
	if isSuccess {
		fmt.Println("Successfully scraping data")
	}
}
