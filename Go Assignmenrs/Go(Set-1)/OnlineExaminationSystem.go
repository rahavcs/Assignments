package main

import "fmt"

type Question struct {
	Question   string
	Options    [4]string
	CorrectAns int
}

func main() {
	// Define the question bank with answers
	questions := []Question{
		{"What is the capital of India?", [4]string{"1. New Delhi", "2. London", "3. Berlin", "4. Madrid"}, 1},
		{"Which planet is known as the Red Planet?", [4]string{"1. Earth", "2. Mars", "3. Venus", "4. Jupiter"}, 2},
		{"Who wrote Romeo and Juliet?", [4]string{"1. William Shakespeare", "2. Charles Dickens", "3. Mark Twain", "4. Jane Austen"}, 1},
	}

	score := 0

	// Loop through each question
	for i, q := range questions {
		// Display question and options
		fmt.Printf("Question %d: %s\n", i+1, q.Question)
		for _, option := range q.Options {
			fmt.Println(option)
		}

		// User input for answer
		var input string
		fmt.Print("Enter your answer (1-4) or type exit to quit: ")
		fmt.Scanln(&input)

		// Exit condition
		if input == "exit" {
			fmt.Println("You chose to exit the quiz.")
			break
		}

		// Check if input matches correct answer
		if input == fmt.Sprintf("%d", q.CorrectAns) {
			fmt.Println("Correct!")
			score+=1
		} else {
			fmt.Println("Incorrect.")
		}
	}

	// Analyzing performance
	fmt.Printf("Your final score: %d/%d\n", score, len(questions))
	if score == len(questions) {
		fmt.Println("Excellent!")
	} else if score >= len(questions)/2 {
		fmt.Println("Good!")
	} else {
		fmt.Println("Needs Improvement.")
	}
}
