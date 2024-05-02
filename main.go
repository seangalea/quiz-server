package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Starting Quiz REST webservice")

	initQuestions()

	router := gin.Default()
	v1 := router.Group("/quiz/api/v1")

	v1.GET("/questions", getQuestions)
	v1.POST("/answers", postAnswers)
	v1.GET("/user/:user/statistics", getStatistics)

	log.Printf("Running router")
	router.Run(addr)
}

func getQuestions(c *gin.Context) {
	log.Print("GET questions")

	c.IndentedJSON(http.StatusOK, questionsList)
}

func getStatistics(c *gin.Context) {
	log.Print("GET statistics")

	userId := c.Param("user")
	stats := calculateStats(userId)
	c.IndentedJSON(http.StatusOK, stats)
}

func postAnswers(c *gin.Context) {
	log.Print("POST answers")

	user := c.Query("user")

	// deserialize JSON
	var answerMatrix = AnswerMatrix{}
	if err := c.BindJSON(&answerMatrix); err != nil {
		return
	}

	// record score
	score := calculateScore(answerMatrix.Answers)
	answerMatrix.Score = score
	recordedScores[user] = score

	// return updated resource
	c.IndentedJSON(http.StatusCreated, answerMatrix)
}
