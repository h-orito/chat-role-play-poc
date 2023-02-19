package model

type GameRepository interface {
	Find(ID uint32) (game *Game, err error)
	Save(game *Game) (saved *Game, err error)
}
