package main

import (
	"log"
	"sort"
)

func initQuestions() {
	log.Print("Initializing questions")

	// generate question list
	for key, value := range questionsRepository {
		value.ID = key
		questionsList = append(questionsList, value)
	}

	// Sort questions array by ID
	sort.Slice(questionsList, func(i, j int) bool {
		if questionsList[i].ID == questionsList[j].ID {
			log.Fatalf("Found duplicate question ID: %d", questionsList[i].ID)
		}
		return questionsList[i].ID < questionsList[j].ID
	})
}

func calculateScore(answers []Answer) int {
	log.Print("Calculating score")

	wasAnsweredCorrectly := make(map[int]bool)
	score := 0
	for _, ans := range answers {
		wasCorrect, wasAnswered := wasAnsweredCorrectly[ans.QuestionID]

		// Revert score if the question has been answered before
		if wasAnswered && wasCorrect {
			score--
		}

		// keep track of answer and increment score only if answer is correct
		if ans.AnswerID == questionsRepository[ans.QuestionID].CorrectAnswer {
			score++
			wasAnsweredCorrectly[ans.QuestionID] = true
		} else {
			wasAnsweredCorrectly[ans.QuestionID] = false
		}
	}

	return score
}

func calculateStats(user string) Statistic {
	log.Printf("Calculating stats for %s", user)

	stat := Statistic{}
	stat.UserScore = recordedScores[user]
	stat.TotalQuizzers = len(recordedScores)

	// count the number of participants scoring lower than this user
	worseQuizzers := 0
	for _, score := range recordedScores {
		if score < stat.UserScore {
			worseQuizzers++
		}
	}

	// calculate the ratio of worseQuizzers : totalQuizzers
	if stat.TotalQuizzers > 0 {
		stat.WorseQuizzersRatio = (float32(worseQuizzers) / float32(stat.TotalQuizzers))
	}

	return stat
}
