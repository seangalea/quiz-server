package main

import (
	"log"
)

func calculateScore(answerList []Answer) int {
	log.Print("Calculating score")

	score := 0
	for i, ans := range answerList {
		// increment score if answer is correct
		if i < len(questions) && ans.AnswerID == questions[i].CorrectAnswer {
			score++
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
