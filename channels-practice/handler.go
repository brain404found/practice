package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func processMathParams(c *gin.Context, workerPool chan Params) {
	startTime := time.Now()

	var requestParams Params
	if err := c.BindJSON(&requestParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	mathParams := generateParams()
	requestParams.MathParams = mathParams
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
