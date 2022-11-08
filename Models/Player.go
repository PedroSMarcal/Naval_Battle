package models

type player struct {
	ID     uint
	Name   string
	Age    uint
	Score  uint
	region string
	rank   uint
	Board  Board
}

type PlayerInterface interface {
	Search(id uint) player
}

func NewPlayer() player {
	return player{}
}
