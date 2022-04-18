package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	response := gin.H{
		"status":  "OK! 🚀",
		"message": "Let's get started to build your awesome apps 🔥🔥🔥",
		"developer": map[string]interface{}{
			"name":     "Sahibul Nuzul Firdaus",
			"email":    "sahibulnuzulfirdaus13@gmail.com",
			"linkedIn": "https://www.linkedin.com/in/sahibul-nf/",
			"support": map[string]string{
				"paypal 💰":               "https://www.paypal.com/paypalme/sahibulnf",
				"buy me a coffee ☕️":     "https://www.buymeacoffee.com/sahibulnf",
				"karyakarsa":             "https://karyakarsa.com/sahibul_nf",
				"send thank you message": "https://wa.link/r3amjb",
			},
		},
		"source": "https://github.com/sahibul-nf/aceh-dictionary-api",
		"endpoints": map[string]interface{}{
			"advices": map[string]string{
				"method":      http.MethodGet,
				"pattern":     "http://aceh-dictionary.herokuapp.com/api/v1/advices?input=YOUR_QUERY",
				"example":     "http://aceh-dictionary.herokuapp.com/api/v1/advices?input=lon",
				"description": "Returns a list of suggested words that match the Input.",
			},
			"dictionary": map[string]string{
				"method":      http.MethodGet,
				"pattern":     "http://aceh-dictionary.herokuapp.com/api/v1/dictionaries",
				"description": "Returns a list of dictionary data.",
			},
		},
	}

	ctx.JSON(http.StatusOK, response)
}
