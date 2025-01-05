package model

type Question struct {
	Question      string
	Options       [4]string
	CorrectAnswer int // Correct answer index (1-based)
}
