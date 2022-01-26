package scraping

import (
	"aceh-dictionary-api/dictionary"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Service interface {
	FetchWordByCharPerPage() []string
	ListOfChar() []string
	CheckPageCount(char string) int
	FetchAcehIndoDictionary() []dictionary.AcehIndo
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func ListOfChar() []string {
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "r", "s", "t", "u", "p", "w", "y", "z"}
	return chars
}

func CheckPageCount(char string) int {
	pageCount := 0

	res, err := http.Get(fmt.Sprintf("https://kata.web.id/kamus/indonesia-aceh/huruf/%s", char))
	if err != nil {
		// log.Fatal(err)
		pageCount = 0
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		fmt.Printf("status code error: %d %s \n", res.StatusCode, res.Status)
		pageCount = 0
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		// log.Fatal(err)
		pageCount = 0
	}

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(j int, t *goquery.Selection) {
			pageCount = j
		})
	})

	if pageCount != 0 {
		pageCount -= 1
	}

	return pageCount
}

func FetchWordByCharPerPage() []string {

	chars := ListOfChar()
	words := []string{}
	var page int

	for _, char := range chars {
		page = CheckPageCount(char)

		if page != 0 {
			for i := 1; i <= page; i++ {
				res, err := http.Get(fmt.Sprintf("https://kata.web.id/kamus/indonesia-aceh/huruf/%s/page/%d", char, i))
				if err != nil {
					log.Fatal(err)
				}

				defer res.Body.Close()

				if res.StatusCode != 200 {
					log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
				}

				doc, err := goquery.NewDocumentFromReader(res.Body)
				if err != nil {
					log.Fatal(err)
				}

				doc.Find(".word-list").Each(func(i int, s *goquery.Selection) {
					s.Find("li").Each(func(j int, t *goquery.Selection) {
						// fmt.Println(t.Text())
						words = append(words, t.Text())
					})
				})

				// fmt.Println()
			}
		} else {
			res, err := http.Get(fmt.Sprintf("https://kata.web.id/kamus/indonesia-aceh/huruf/%s", char))
			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()

			if res.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}

			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			doc.Find(".word-list").Each(func(i int, s *goquery.Selection) {
				s.Find("li").Each(func(j int, t *goquery.Selection) {
					// fmt.Println(t.Text())
					words = append(words, t.Text())
				})
			})

			// fmt.Println()
		}

	}

	return words
}

func FetchAcehIndoDictionary() []dictionary.DictionaryInput {
	rows := []dictionary.DictionaryInput{}

	words := FetchWordByCharPerPage()
	fmt.Println(len(words))
	// lengthWord := len(words)

	for _, word := range words {
		fmt.Println(word)
		fmt.Println()
		// process := float64((i / lengthWord) * (100 / 100))
		// fmt.Println(process)

		newWord := word

		if strings.Contains(word, " ") {
			newWord = strings.Replace(word, " ", "-", -1)
			// fmt.Println(newWord)
		}

		res, err := http.Get(fmt.Sprintf("https://kata.web.id/kamus/indonesia-aceh/arti-kata/%s", newWord))
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
			row := dictionary.DictionaryInput{}
			var aceh []string
			var indo string

			s.Find("td").Each(func(j int, t *goquery.Selection) {

				if j > 0 {
					aceh = strings.Split(t.Text(), ", ")
				} else {
					indo = t.Text()
				}

				for _, v := range aceh {
					row.Indonesia = indo
					row.Aceh = v
					rows = append(rows, row)
				}
			})
		})

	}

	// fmt.Println(len(rows))

	return rows
}
