package main

import (
	"aceh-dictionary-api/config"
	"aceh-dictionary-api/dictionary"
	"aceh-dictionary-api/search"
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var ()

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := config.SetupDatabaseConnection()

	searchRepository := search.NewRepository(db)
	searchService := search.NewService(searchRepository)

	defer config.CloseDatabaseConnection(db)

	// ô, è, é, ö
	// Buat regular expression yang mencocokkan string yang memiliki tanda è, é, atau ô
	re := []*regexp.Regexp{
		regexp.MustCompile("è"),
		regexp.MustCompile("é"),
		regexp.MustCompile("ô"),
	}

	scanner := bufio.NewScanner(os.Stdin)
	// masukkan jumlah data yang ingin diambil
	fmt.Print("Masukkan jumlah data sample yang ingin diambil: ")
	scanner.Scan()
	amount, _ := strconv.Atoi(scanner.Text())

	// masukkan algoritma yang ingin digunakan
	fmt.Print("Masukkan algoritma yang ingin digunakan (lev atau jwd): ")
	scanner.Scan()
	algorithm := strings.ToLower(scanner.Text())

	// ambil data sample dari database
	fmt.Println("Mengambil data sample dari database...")
	samples := GetSampleData(db, re, amount)
	sampleKeywords, expectedKeywords := GetKeywordsAndExpectedKeyword(samples, re)
	// randomKeywords := GetRandomKeywords(amount, sampleKeywords)

	listOfAccuracy := []Accuracy{}

	for i, v := range sampleKeywords {
		fmt.Printf("\nKeyword ke-%d: %s\n", i, v)

		result, _ := searchService.GetRecommendationWords(v, algorithm)
		order := 0
		expectedKeyword := ""
		accuracy := Accuracy{}

		for i, r := range result {
			order = i + 1
			fmt.Printf("%d. %s = %.2f\n", order, r.Aceh, r.Similiarity)

			expectedKeyword = GetExpectedResult(r.ID, expectedKeywords)
			if expectedKeyword != "" {
				value := AccuracyCalculation(order, len(result))
				accuracy = Accuracy{
					keyword:              v,
					expected:             expectedKeyword,
					algorithmResult:      r.Similiarity,
					order:                order,
					recommendationResult: value,
				}
				break
			}
		}

		listOfAccuracy = append(listOfAccuracy, accuracy)

		fmt.Println("=====================================")
	}

	finalResult := AccuracyPercentageCalculation(listOfAccuracy, len(sampleKeywords))
	fmt.Printf("\nTotal percentage of accuracy for %d data sample using %s algorithm is %.1f%%\n", len(sampleKeywords), algorithm, finalResult.recommendationAccuracyPercent)

	// save to csv
	fmt.Println("\nMenyimpan data ke file csv...")
	SaveToCSV(listOfAccuracy, algorithm, finalResult)

	fmt.Println("Selesai 🤩")
}

func GetSampleData(db *gorm.DB, re []*regexp.Regexp, count int) []dictionary.Dictionary {
	// Cari data string yang memiliki tanda è, é, atau ô di database
	var data []dictionary.Dictionary
	db.Where("aceh ~* ?", re[0].String()).Or("aceh ~* ?", re[1].String()).Or("aceh ~* ?", re[2].String()).Find(&data)

	sample := []dictionary.Dictionary{}
	// Gunakan data yang ditemukan sesuai kebutuhan Anda
	for _, d := range data {
		// skip jika data lebih dari 1 kata atau mengandung spasi
		if strings.Contains(d.Aceh, " ") {
			continue
		}

		if len(sample) == count {
			break
		}

		// fmt.Printf("ID: %d, WORD: %s\n", d.ID, d.Aceh)
		sample = append(sample, d)
	}

	fmt.Println("Len: ", len(sample))

	return sample
}

// func for get keyword
func GetKeywordsAndExpectedKeyword(data []dictionary.Dictionary, re []*regexp.Regexp) ([]string, []dictionary.Dictionary) {

	keywords := []string{}
	expectedKeywords := []dictionary.Dictionary{}

	// hapus tanda è, é, atau ô
	for _, d := range data {
		expectedKeywords = append(expectedKeywords, d)

		for _, r := range re {
			if r.String() == "è" || r.String() == "é" {
				d.Aceh = r.ReplaceAllString(d.Aceh, "e")
			}

			if r.String() == "ô" {
				d.Aceh = r.ReplaceAllString(d.Aceh, "o")
			}
		}

		// fmt.Printf("ID: %d, WORD: %s\n", d.ID, d.Aceh)
		keywords = append(keywords, d.Aceh)
	}

	return keywords, expectedKeywords
}

func GetRandomKeywords(amount int, keywords []string) []string {
	var result []string

	// Shuffle keyword array
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keywords), func(i, j int) {
		keywords[i], keywords[j] = keywords[j], keywords[i]
	})

	// Ambil `amount` keyword pertama
	result = keywords[:amount]

	return result
}

// masukkan hasil ke dalam file csv
func SaveToCSV(data []Accuracy, algorithm string, finalResult FinalResult) {
	// file name
	fileName := fmt.Sprintf("result-%s-%d.csv", algorithm, len(data))

	// Buat file CSV
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Buat writer CSV
	writer := csv.NewWriter(file)
	// Flush writer CSV
	writer.Flush()

	// Tulis data ke file CSV
	var rows [][]string
	rows = append(rows, []string{"Keyword", "Expected", "Algorithm Accuracy Result", "Priority Number", "Recommendation Accuracy Result"})

	for _, v := range data {
		row := []string{
			v.keyword,
			v.expected,
			strconv.FormatFloat(v.algorithmResult, 'f', 2, 64),
			strconv.Itoa(v.order),
			strconv.FormatFloat(v.recommendationResult, 'f', 2, 64),
		}
		rows = append(rows, row)
	}

	// add total result
	rows = append(rows, []string{"Total Accuracy Results:", "", strconv.FormatFloat(finalResult.algorithm, 'f', 2, 64), "", strconv.FormatFloat(finalResult.recommendation, 'f', 2, 64)})
	// add total accuracy
	rows = append(rows, []string{"Total Percentage Accuracy:", "", strconv.FormatFloat(finalResult.algorithmAccuracyPercent, 'f', 1, 64) + "%", "", strconv.FormatFloat(finalResult.recommendationAccuracyPercent, 'f', 1, 64) + "%"})

	err = writer.WriteAll(rows)
	if err != nil {
		log.Fatal(err)
	}
}

// accuracy calculation based on the priority number of recommendations
func AccuracyCalculation(number int, amountOfRecommendationResult int) float64 {
	priorityNumber := PriorityNumberComparison(number, amountOfRecommendationResult)

	result := float64(priorityNumber) / float64(amountOfRecommendationResult)
	fmt.Printf("Priority Number: %d / amountOfRecommendationResult: %d = %.1f\n", priorityNumber, amountOfRecommendationResult, result)

	return result
}

// accuracy percentage calculation
func AccuracyPercentageCalculation(listOfAccuracy []Accuracy, amountOfSample int) FinalResult {
	var recomResult float64
	var totalRecommendationAcc float64
	var totalAlgorithmAcc float64

	for _, v := range listOfAccuracy {
		totalRecommendationAcc += v.recommendationResult
		totalAlgorithmAcc += v.algorithmResult
		// fmt.Printf("Keyword: %s\n", v.keyword)
		// fmt.Printf("Expected: %s\n", v.expected)
		// fmt.Printf("Algorithm Result: %.1f\n", v.algorithmResult)
		// fmt.Printf("Order: %d\n", v.order)
		// fmt.Printf("Accuracy: %.1f\n", v.value)

		// fmt.Printf("============\n\n")
	}

	recomResult = totalRecommendationAcc / float64(amountOfSample) * 100
	algoResult := totalAlgorithmAcc / float64(amountOfSample) * 100

	finalResult := FinalResult{
		recommendationAccuracyPercent: recomResult,
		recommendation:                totalRecommendationAcc,
		algorithmAccuracyPercent:      algoResult,
		algorithm:                     totalAlgorithmAcc,
	}

	return finalResult
}

// priority number comparison
func PriorityNumberComparison(number int, amountOfRecommendationResult int) int {
	threshhold := amountOfRecommendationResult

	if threshhold == 10 || threshhold == 5 {
		if number > threshhold {
			return 0
		}
	}

	if number == 1 {
		return amountOfRecommendationResult
	}

	return amountOfRecommendationResult - number + 1
}

func GetExpectedResult(wordId int, expectedKeywords []dictionary.Dictionary) string {

	var expextedResult string
	for _, v := range expectedKeywords {
		if v.ID == wordId {
			expextedResult = v.Aceh
			break
		}
	}

	return expextedResult
}

type Accuracy struct {
	keyword              string
	expected             string
	order                int
	recommendationResult float64
	algorithmResult      float64
}

type FinalResult struct {
	recommendationAccuracyPercent float64
	recommendation                float64
	algorithmAccuracyPercent      float64
	algorithm                     float64
}
