package entity

type Question struct {
	Id              uint
	Text            string
	PossibleAnswers []PossibleAnswer
	CorrectAnswerId uint
	CategoryId      uint
	Difficulty      string
}

type QuestionDifficulty uint8

const (
	Easy QuestionDifficulty = iota + 1
	Medium
	Hard
)

func (q QuestionDifficulty) IsValid() bool {
	if q < Easy || q > Hard {
		return false
	}

	return true

}

type PossibleAnswer struct {
	Id     uint
	Text   string
	Choice PossibleAnswerChoice
}

type PossibleAnswerChoice uint8

func (p PossibleAnswerChoice) IsValid() bool {
	if p < PossibleAnswerA || p > PossibleAnswerD {
		return false
	}

	return true
}

const (
	PossibleAnswerA PossibleAnswerChoice = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)
