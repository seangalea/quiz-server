package main

const addr = "localhost:8080"

var recordedScores = make(map[string]int)
var questions = []Question{
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
	{
		ID:            3,
		Query:         "What is the chemical symbol for water?",
		Answers:       []string{"C02", "H2O", "NaCl", "O2"},
		CorrectAnswer: 1,
	},
	{
		ID:            4,
		Query:         "What is the largest mammal?",
		Answers:       []string{"Elephant", "Blue Whale", "Giraffe", "Horse"},
		CorrectAnswer: 1,
	},
	{
		ID:            5,
		Query:         "Which is the tallest mountain in the world?",
		Answers:       []string{"K2", "Mount Everest", "Kangchenjunga", "Lhotse"},
		CorrectAnswer: 1,
	},
	{
		ID:            6,
		Query:         "What is the chemical symbol for gold?",
		Answers:       []string{"Au", "Ag", "Fe", "Cu"},
		CorrectAnswer: 0,
	},
	{
		ID:            7,
		Query:         "Which animal is known as the 'king of the jungle'?",
		Answers:       []string{"Lion", "Tiger", "Elephant", "Giraffe"},
		CorrectAnswer: 0,
	},
	{
		ID:            8,
		Query:         "What is the hardest natural substance on Earth?",
		Answers:       []string{"Diamond", "Steel", "Quartz", "Graphite"},
		CorrectAnswer: 0,
	},
	{
		ID:            9,
		Query:         "What is the largest organ in the human body?",
		Answers:       []string{"Brain", "Skin", "Heart", "Liver"},
		CorrectAnswer: 1,
	},
	{
		ID:            10,
		Query:         "Which gas do plants absorb for photosynthesis?",
		Answers:       []string{"Oxygen", "Carbon Dioxide", "Nitrogen", "Hydrogen"},
		CorrectAnswer: 1,
	},
	{
		ID:            11,
		Query:         "Who wrote 'Romeo and Juliet'?",
		Answers:       []string{"William Shakespeare", "Jane Austen", "Charles Dickens", "Mark Twain"},
		CorrectAnswer: 0,
	},
	{
		ID:            12,
		Query:         "What is the chemical symbol for oxygen?",
		Answers:       []string{"O2", "O", "O3", "CO2"},
		CorrectAnswer: 1,
	},
	{
		ID:            13,
		Query:         "What is the primary function of the lungs?",
		Answers:       []string{"Pumping blood", "Digesting food", "Producing hormones", "Breathing"},
		CorrectAnswer: 3,
	},
	{
		ID:            14,
		Query:         "Which country is known as the 'Land of the Rising Sun'?",
		Answers:       []string{"China", "India", "Japan", "South Korea"},
		CorrectAnswer: 2,
	},
	{
		ID:            15,
		Query:         "What is the chemical symbol for sodium?",
		Answers:       []string{"Na", "Sn", "Fe", "Si"},
		CorrectAnswer: 0,
	},
	{
		ID:            16,
		Query:         "What is the process of plants making their own food called?",
		Answers:       []string{"Photosynthesis", "Respiration", "Fermentation", "Pollination"},
		CorrectAnswer: 0,
	},
	{
		ID:            17,
		Query:         "Which element is a gas at room temperature?",
		Answers:       []string{"Iron", "Mercury", "Oxygen", "Sulfur"},
		CorrectAnswer: 2,
	},
	{
		ID:            18,
		Query:         "What is the currency of Japan?",
		Answers:       []string{"Dollar", "Yen", "Euro", "Pound"},
		CorrectAnswer: 1,
	},
	{
		ID:            19,
		Query:         "Which of the following is a prime number?",
		Answers:       []string{"12", "15", "17", "20"},
		CorrectAnswer: 2,
	},
	{
		ID:            20,
		Query:         "What is the chemical symbol for carbon?",
		Answers:       []string{"Ca", "Cr", "C", "Co"},
		CorrectAnswer: 2,
	},
}
