package handler

import (
	"fmt"
	"online-exam/service"
)

func StartQuiz() {
	quizService := service.NewQuizService()

	fmt.Println("Welcome to the Online Quiz!")
	err := quizService.TakeQuiz()
	if err != nil {
		fmt.Println(err)
		return
	}

	performance := quizService.CalculatePerformance()
	fmt.Println("Your performance: ", performance)
}
