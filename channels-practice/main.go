package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Params struct {
	UserID     string   `json:"user_id"`
	MathParams []string `json:"math_params"`
}

type MathParam struct {
	UserID string `json:"user_id"`
	Param  string `json:"-"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "postgres://postgres:postgres@localhost:5433/postgres",
	}))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&MathParam{})

	router := gin.Default()
	router.POST("/ProcessMathParams", processMathParams)
	router.Run(":8080")
}

func generateParams() []string {
    var result []string
    for i := 0; i < 1000; i++ {
        result = append(result, fmt.Sprintf("param: %d", i))
    }
    return result
}

func processMathParams(c *gin.Context) {
    startTime := time.Now()

    mathParams := generateParams()

    var requestParams Params
    requestParams.MathParams = mathParams
    if err := c.BindJSON(&requestParams); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    paramChannel := make(chan string, 100)
    poolSize := 10

    var workerWaitGroup sync.WaitGroup
    workerWaitGroup.Add(poolSize)

    for i := 0; i < poolSize; i++ {
        go processParams(requestParams, paramChannel, &workerWaitGroup)
    }

    go writeToDB(paramChannel)

    workerWaitGroup.Wait()

    close(paramChannel)

    elapsed := time.Since(startTime)

    c.JSON(200, gin.H{"message": fmt.Sprintf("Processed math params in %v", elapsed)})
}


func processParams(params Params, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(params.MathParams); i += 100 {
		end := i + 100
		if end > len(params.MathParams) {
			end = len(params.MathParams)
		}
		for _, p := range params.MathParams[i:end] {
			time.Sleep(100 * time.Millisecond)
			ch <- p
		}
	}
}

func writeToDB(ch chan string) {
	counter := 0
	var batch []string
	for params := range ch {
		batch = append(batch, params)
		counter++
		if counter == 100 {
			writeBatchToDB(batch)
			counter = 0
			batch = nil
		}
	}

	if len(batch) > 0 {
		writeBatchToDB(batch)
	}
}

func writeBatchToDB(params []string) {
	var mathParams []MathParam
	for _, p := range params {
		mathParams = append(mathParams, MathParam{Param: p})
	}

	if err := db.Create(&mathParams).Error; err != nil {
		log.Printf("Error writing to DB: %s\n", err.Error())
	}
}

