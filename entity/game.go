package entity

type Game struct {
	Id          uint
	CategoryId  uint
	QuestionIds []uint
	Players     []Player
}

type Player struct {
	Id      uint
	UserId  uint
	GameId  uint
	Score   uint
	Answers []PlayerAnswer
}

type PlayerAnswer struct {
	Id         uint
	PlayerId   uint
	QuestionId uint
}
