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

	workerPool := createWorkerPool(10)

	router := gin.Default()
	router.POST("/ProcessMathParams", func(c *gin.Context) {
		processMathParams(c, workerPool)
	})
	router.Run(":8080")
}

func processMathParams(c *gin.Context, workerPool chan Params) {
	startTime := time.Now()

	var requestParams Params
	if err := c.BindJSON(&requestParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	requestParams.MathParams = generateParams()
	workerPool <- requestParams

	elapsed := time.Since(startTime)
	c.JSON(200, gin.H{"message": fmt.Sprintf("Processed math params in %v", elapsed)})
}

func generateParams() []string {
	var result []string
	for i := 0; i < 1000; i++ {
		result = append(result, fmt.Sprintf("param: %d", i))
	}
	return result
}

func createWorkerPool(poolSize int) chan Params {
	paramCh := make(chan Params) // Используем небуферизированный канал
	var wgWorker sync.WaitGroup

	for i := 0; i < poolSize; i++ {
		wgWorker.Add(1)
		go func() {
			defer wgWorker.Done()
			for params := range paramCh {
				// Обработка запросов батчем
				processParams(params)
			}
		}()
	}

	go func() {
		wgWorker.Wait()
		close(paramCh)
	}()

	return paramCh
}

func processParams(params Params) {
	// Имитация деятельности на 0.1 секунды
	time.Sleep(100 * time.Millisecond)

	var mathParamsToWrite []MathParam
	for _, p := range params.MathParams {
		mathParamsToWrite = append(mathParamsToWrite, MathParam{UserID: params.UserID, Param: p})
	}

	if err := writeToDB(mathParamsToWrite); err != nil {
		log.Printf("Error writing to DB: %s\n", err.Error())
	}
}

func writeToDB(params []MathParam) error {
	return db.Create(&params).Error
}
