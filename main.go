package main

import (
	"aceh-dictionary-api/config"
	"aceh-dictionary-api/dictionary"
	"aceh-dictionary-api/search"
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
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

	// masukkan nilai relevansi yang ingin digunakan
	fmt.Print("Masukkan nilai relevansi yang ingin digunakan (0.0 - 1.0): ")
	scanner.Scan()
	threshold, _ := strconv.ParseFloat(scanner.Text(), 64)

	// masukkan metode akurasi yang ingin digunakan (map atau default)
	fmt.Print("Masukkan metode akurasi yang ingin digunakan, default(0) atau map(1): ")
	scanner.Scan()
	method, _ := strconv.Atoi(scanner.Text())

	// ambil data sample dari database
	fmt.Println("Mengambil data sample dari database...")
	samples := GetSampleData(db, re, amount)
	sampleKeywords, expectedKeywords := GetKeywordsAndExpectedKeyword(samples, re)
	// randomKeywords := GetRandomKeywords(amount, sampleKeywords)

	listOfAccuracy := []Accuracy{}

	for j, v := range sampleKeywords {
		fmt.Printf("\nKeyword ke-%d: %s\n", j, v)

		result, _ := searchService.GetRecommendationWords(v, algorithm)
		relList := FilterByThreshold(result, threshold)

		priorityNumber := 0
		expectedKeyword := ""
		accuracy := Accuracy{}

		if method == 1 {
			ap := CalculateAveragePrecision(relList, len(result))
			accuracy.ap = ap
			accuracy.keyword = v
		}

		for i, r := range result {
			priorityNumber = i + 1
			fmt.Printf("%d. %s = %.2f\n", priorityNumber, r.Aceh, r.Similiarity)

			expectedKeyword = GetExpectedResult(r.ID, expectedKeywords)
			if expectedKeyword == "" {
				expectedKeyword = expectedKeywords[j].Aceh
			}

			if r.Aceh == expectedKeyword {
				if method == 1 {
					accuracy.expected = expectedKeyword
					accuracy.algorithmResult = r.Similiarity
					accuracy.priorityNumber = priorityNumber
				} else {
					value := AccuracyCalculation(priorityNumber, len(result))
					accuracy = Accuracy{
						keyword:         v,
						expected:        expectedKeyword,
						algorithmResult: r.Similiarity,
						priorityNumber:  priorityNumber,
						value:           value,
					}
				}
				break
			} else {
				accuracy.expected = expectedKeyword
				accuracy.algorithmResult = 0
				accuracy.priorityNumber = 0
			}
		}

		listOfAccuracy = append(listOfAccuracy, accuracy)

		fmt.Println("=====================================")
	}

	var finalResult FinalResult
	var mapResult MAPResult

	if method == 1 {
		mapResult = CalculateMAP(listOfAccuracy)
		fmt.Printf("\nTotal percentage of accuracy for %d data sample with threshold %.1f using %s algorithm is %.1f%%\n", len(sampleKeywords), threshold, algorithm, mapResult.percentage)
	} else {
		finalResult = AccuracyPercentageCalculation(listOfAccuracy, len(sampleKeywords))
		fmt.Printf("\nTotal percentage of accuracy for %d data sample using %s algorithm is %.1f%%\n", len(sampleKeywords), algorithm, finalResult.RecommendationAccuracyPercent)
	}

	// save to csv
	fmt.Println("\nMenyimpan data ke file csv...")
	SaveToCSV(listOfAccuracy, algorithm, finalResult, mapResult, method, threshold)

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
func SaveToCSV(data []Accuracy, algorithm string, finalResult FinalResult, mapResult MAPResult, method int, threshold float64) {
	// file name
	fileName := fmt.Sprintf("result-%s-%d.csv", algorithm, len(data))
	if method == 1 {
		fileName = fmt.Sprintf("result-map-%s-%d-%.1f.csv", algorithm, len(data), threshold)
	}

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
	if method == 1 {
		// rows = append(rows, []string{"No", "Keyword", "Expectation", "Algorithm Accuracy Result", "Priority Number", "Precision", "Recall", "Average Precision"})
		rows = append(rows, []string{"No", "Keyword", "Expectation", "Algorithm Accuracy Result", "Priority Number", "Average Precision"})
	} else {
		rows = append(rows, []string{"No", "Keyword", "Expectation", "Algorithm Accuracy Result", "Priority Number", "Recommendation Accuracy Result"})
	}

	for i, v := range data {
		var row []string
		number := strconv.Itoa(i + 1)

		if method == 1 {
			row = []string{
				number,
				v.keyword,
				v.expected,
				strconv.FormatFloat(v.algorithmResult, 'f', 2, 64),
				strconv.Itoa(v.priorityNumber),
				// strconv.FormatFloat(v.ap.Precision, 'f', 2, 64),
				// strconv.FormatFloat(v.ap.Recall, 'f', 2, 64),
				strconv.FormatFloat(v.ap.AveragePrecision, 'f', 2, 64),
			}
		} else {
			row = []string{
				number,
				v.keyword,
				v.expected,
				strconv.FormatFloat(v.algorithmResult, 'f', 2, 64),
				strconv.Itoa(v.priorityNumber),
				strconv.FormatFloat(v.value, 'f', 2, 64),
			}
		}
		rows = append(rows, row)
	}

	if method == 1 {
		// add total result
		// rows = append(rows, []string{"", "MAP Results:", "", "", "", "", "", strconv.FormatFloat(mapResult.value, 'f', 2, 64)})
		rows = append(rows, []string{"", "MAP Results:", "", "", "", strconv.FormatFloat(mapResult.value, 'f', 2, 64)})
		// add total accuracy
		// rows = append(rows, []string{"", "MAP Percentage Accuracy:", "", "", "", "", "", strconv.FormatFloat(mapResult.percentage, 'f', 1, 64) + "%"})
		rows = append(rows, []string{"", "MAP Percentage Accuracy:", "", "", "", strconv.FormatFloat(mapResult.percentage, 'f', 1, 64) + "%"})
	} else {
		// add total result
		rows = append(rows, []string{"", "Total Accuracy Results:", "", strconv.FormatFloat(finalResult.Algorithm, 'f', 2, 64), "", strconv.FormatFloat(finalResult.Recommendation, 'f', 2, 64)})
		// add total accuracy
		rows = append(rows, []string{"", "Total Percentage Accuracy:", "", strconv.FormatFloat(finalResult.AlgorithmAccuracyPercent, 'f', 1, 64) + "%", "", strconv.FormatFloat(finalResult.RecommendationAccuracyPercent, 'f', 1, 64) + "%"})

	}

	err = writer.WriteAll(rows)
	if err != nil {
		log.Fatal(err)
	}
}

// accuracy calculation based on the priority number of recommendations
func AccuracyCalculation(number int, amountOfRecommendationResult int) float64 {
	priorityNumber := PriorityNumberComparison(number, amountOfRecommendationResult)

	result := float64(priorityNumber) / float64(amountOfRecommendationResult)
	// fmt.Printf("Priority Number: %d / amountOfRecommendationResult: %d = %.1f\n", priorityNumber, amountOfRecommendationResult, result)

	return result
}

// accuracy percentage calculation
func AccuracyPercentageCalculation(listOfAccuracy []Accuracy, amountOfSample int) FinalResult {
	var recomResult float64
	var totalRecommendationAcc float64
	var totalAlgorithmAcc float64

	for _, v := range listOfAccuracy {
		totalRecommendationAcc += v.value
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
		RecommendationAccuracyPercent: recomResult,
		Recommendation:                totalRecommendationAcc,
		AlgorithmAccuracyPercent:      algoResult,
		Algorithm:                     totalAlgorithmAcc,
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
	keyword         string
	expected        string
	priorityNumber  int
	algorithmResult float64
	value           float64
	ap              APResult
}

type APResult struct {
	Precision        float64
	Recall           float64
	AveragePrecision float64
}

type FinalResult struct {
	RecommendationAccuracyPercent float64
	Recommendation                float64
	AlgorithmAccuracyPercent      float64
	Algorithm                     float64
}

type MAPResult struct {
	value      float64
	percentage float64
}

type RelevantList struct {
	Keyword string
	Label   float64 // 1 = relevant, 0 = not relevant
}

// calculate the accuracy with the AP method (average precision) and return the average precision
// Menghitung Average Precision (AP)
// k = jumlah data yang akan dihitung dari hasil rekomendasi (misal: 10 data pertama atau 5 data pertama)
func CalculateAveragePrecision(relList []RelevantList, k int) APResult {
	// Menghitung Precision@k
	// precisionAtK := float64(0)
	// for i := 0; i < k; i++ {
	// 	if relList[i].Label == 1 {
	// 		precisionAtK++
	// 	}
	// }
	// precisionAtK /= float64(k)

	// // Menghitung recall
	// recall := float64(0)
	// for _, rel := range relList {
	// 	if rel.Label == 1 {
	// 		recall++
	// 	}
	// }
	// recall /= float64(len(relList))

	R := float64(0)            // jumlah relevan (true positive / TP = 1)
	precisionAtK := float64(0) // jumlah data yang relevan (true positive / TP = 1) yang ditemukan pada k data pertama

	for i := 0; i < k; i++ {
		if relList[i].Label == 1 {
			R++
			precisionAtK += R / float64(i+1)
			// precisionAtK += 1
		}
	}

	// Menghitung AP
	AP := (1 / R) * precisionAtK
	// AP := R / float64(k)

	if math.IsNaN(AP) {
		AP = 0
	}

	return APResult{
		// Precision:        precision,
		// Recall:           recall,
		AveragePrecision: AP,
	}
}

// calculate MAP (mean average precision) and return the result
func CalculateMAP(listOfAccuracy []Accuracy) MAPResult {
	var mapResult float64

	for _, v := range listOfAccuracy {
		mapResult += v.ap.AveragePrecision
	}

	mapResult /= float64(len(listOfAccuracy))

	// ubah ke persentase
	persentage := mapResult * 100

	return MAPResult{
		value:      mapResult,
		percentage: persentage,
	}
}

// filter the recommendation result based on the threshold value and return the relevant list (keyword and label (1 = relevant, 0 = not relevant))
func FilterByThreshold(recommendationResult []search.RecommendationWord, threshold float64) []RelevantList {
	var result []RelevantList
	for _, v := range recommendationResult {
		if v.Similiarity >= threshold {
			result = append(result, RelevantList{
				Keyword: v.Aceh,
				Label:   1,
			})
		} else {
			result = append(result, RelevantList{
				Keyword: v.Aceh,
				Label:   0,
			})
		}
	}

	return result
}
