package main

import (
	"log"
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
