package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Starting Quiz REST webservice")

	router := gin.Default()
	v1 := router.Group("/quiz/api/v1")

	v1.GET("/questions", getQuestions)
	v1.POST("/answers", postAnswers)
	v1.GET("/user/:user/statistics", getStatistics)

	log.Printf("Running router")
	router.Run(addr)
}
