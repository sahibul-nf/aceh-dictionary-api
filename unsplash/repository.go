package unsplash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Repository interface {
	GetPhotoByKeyword(keyword string, count int, orientation string) (UnsplashResponse, error)
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) GetPhotoByKeyword(keyword string, count int, orientation string) (UnsplashResponse, error) {
	accessKey := os.Getenv("UNSPLASH_ACCESS_KEY")

	// http client to call unsplash api
	client := http.Client{}

	// request to unsplash api
	url := fmt.Sprintf("https://api.unsplash.com/photos/random?query=%s&count=%d&orientation=%s&client_id=%s", keyword, count, orientation, accessKey)
	fmt.Println(url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return UnsplashResponse{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return UnsplashResponse{}, err
	}

	defer res.Body.Close()

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return UnsplashResponse{}, err
	}

	if res.StatusCode != 200 {
		return UnsplashResponse{}, fmt.Errorf("error: %s", resData)
	}

	// handle response
	var response UnsplashResponse
	err = json.Unmarshal(resData, &response)
	if err != nil {
		return UnsplashResponse{}, err
	}

	// return response
	return response, nil
}
