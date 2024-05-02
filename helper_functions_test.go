package main

import (
	"testing"
)

func TestCalculateScore(t *testing.T) {
	// Mock questions
	questionsList = []Question{
		{
			ID:            1,
			Query:         "What is the capital of France?",
			Answers:       []string{"Paris", "London", "Berlin", "Rome"},
			CorrectAnswer: 0,
		},
		{
			ID:            2,
			Query:         "Which planet is known as the Red Planet?",
			Answers:       []string{"Saturn", "Venus", "Jupiter", "Mars"},
			CorrectAnswer: 3,
		},
	}

	// Test cases
	tests := []struct {
		name          string
		answerList    []Answer
		expectedScore int
	}{
		{
			name: "All correct answers",
			answerList: []Answer{
				{QuestionID: 1, AnswerID: 0}, // Correct answer for question 1
				{QuestionID: 2, AnswerID: 3}, // Correct answer for question 2
			},
			expectedScore: 2,
		},
		{
			name: "All incorrect answers",
			answerList: []Answer{
				{QuestionID: 1, AnswerID: 1}, // Incorrect answer for question 1
				{QuestionID: 2, AnswerID: 0}, // Incorrect answer for question 2
			},
			expectedScore: 0,
		},
		{
			name: "Correct and incorrect answers",
			answerList: []Answer{
				{QuestionID: 1, AnswerID: 0}, // Correct answer for question 1
				{QuestionID: 2, AnswerID: 0}, // Incorrect answer for question 2
			},
			expectedScore: 1,
		},
		{
			name: "Multiple answers for the same question",
			answerList: []Answer{
				{QuestionID: 1, AnswerID: 0}, // Ignored answer for question 1
				{QuestionID: 1, AnswerID: 1}, // Ignored answer for question 1
				{QuestionID: 1, AnswerID: 2}, // Ignored answer for question 1
				{QuestionID: 1, AnswerID: 3}, // Incorrect answer for question 1
				{QuestionID: 2, AnswerID: 3}, // Correct answer for question 2
			},
			expectedScore: 1,
		},
		{
			name: "Unordered answers",
			answerList: []Answer{
				{QuestionID: 2, AnswerID: 3},
				{QuestionID: 1, AnswerID: 0},
			},
			expectedScore: 2,
		},
	}

	// Run tests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			score := calculateScore(tc.answerList)
			if score != tc.expectedScore {
				t.Errorf("calculateScore(%v) = %d; expected %d", tc.answerList, score, tc.expectedScore)
			}
		})
	}
}

func TestCalculateStats(t *testing.T) {
	// Set up recorded scores
	recordedScores = map[string]int{
		"User1": 5,
		"User2": 7,
		"User3": 3,
		"User4": 1,
	}

	// Test cases
	testCases := []struct {
		name                       string
		user                       string
		expectedUserScore          int
		expectedWorseQuizzersRatio float32
		expectedTotalQuizzers      int
	}{
		{
			name:                       "Best User",
			user:                       "User2",
			expectedUserScore:          7,
			expectedWorseQuizzersRatio: 0.75,
			expectedTotalQuizzers:      4,
		},
		{
			name:                       "Average User",
			user:                       "User1",
			expectedUserScore:          5,
			expectedWorseQuizzersRatio: 0.5,
			expectedTotalQuizzers:      4,
		},
		{
			name:                       "Worst User",
			user:                       "User4",
			expectedUserScore:          1,
			expectedWorseQuizzersRatio: 0,
			expectedTotalQuizzers:      4,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stat := calculateStats(tc.user)
			if stat.UserScore != tc.expectedUserScore || stat.WorseQuizzersRatio != tc.expectedWorseQuizzersRatio || stat.TotalQuizzers != tc.expectedTotalQuizzers {
				t.Errorf("calculateStats(%s) = %v; expected UserScore: %d, WorseQuizzersRatio: %f, TotalQuizzers: %d",
					tc.user, stat, tc.expectedUserScore, tc.expectedWorseQuizzersRatio, tc.expectedTotalQuizzers)
			}
		})
	}
}
