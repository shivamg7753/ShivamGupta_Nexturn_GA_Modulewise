package service

import (
	"errors"
	"fmt"
	"time"
	"online-exam/model"
)

type QuizService struct {
	Questions []model.Question
	Score     int
}

func NewQuizService() *QuizService {
	return &QuizService{
		Questions: []model.Question{
			{"What is 2 + 2?", [4]string{"3", "4", "5", "6"}, 2},
			{"What is the capital of France?", [4]string{"Berlin", "Madrid", "Paris", "Rome"}, 3},
			{"Go is a ___ programming language.", [4]string{"Functional", "Object-Oriented", "Procedural", "All of the above"}, 4},
		},
	}
}

func (qs *QuizService) TakeQuiz() error {
	for i, question := range qs.Questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, question.Question)
		for j, option := range question.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}

		// Input with validation
		var choice int
		timer := time.NewTimer(15 * time.Second)
		fmt.Println("Enter your answer (1-4) or '0' to exit:")
		
		// Channel for input
		answerChan := make(chan int)
		go func() {
			fmt.Scan(&choice)
			answerChan <- choice
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up! Moving to the next question.")
			continue
		case choice = <-answerChan:
			if choice == 0 {
				return errors.New("Quiz exited by user")
			}
			if choice < 1 || choice > 4 {
				fmt.Println("Invalid choice. Try again.")
				continue
			}

			// Check answer
			if choice == question.CorrectAnswer {
				qs.Score++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Wrong! The correct answer was:", question.Options[question.CorrectAnswer-1])
			}
		}
	}
	return nil
}

func (qs *QuizService) CalculatePerformance() string {
	fmt.Printf("\nYour total score is: %d/%d\n", qs.Score, len(qs.Questions))
	percentage := (float64(qs.Score) / float64(len(qs.Questions))) * 100
	switch {
	case percentage >= 80:
		return "Excellent"
	case percentage >= 50:
		return "Good"
	default:
		return "Needs Improvement"
	}
}
