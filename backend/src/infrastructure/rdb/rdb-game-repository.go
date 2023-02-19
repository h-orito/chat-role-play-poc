package db

import (
	model "chat-role-play-poc/src/domain/model/game"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type GameRepository struct {
	db *DB
}

func NewGameRepository(db *DB) model.GameRepository {
	return &GameRepository{
		db: db,
	}
}

type Game struct {
	ID        uint32
	GameName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (repo *GameRepository) Find(ID uint32) (_ *model.Game, err error) {
	var rdbGame Game
	result := repo.db.Connection.First(&rdbGame, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find: %s \n", err)
	}
	return &model.Game{
		ID:           rdbGame.ID,
		Name:         rdbGame.GameName,
		Participants: model.Participants{},
	}, nil
}

func (repo *GameRepository) Save(g *model.Game) (_ *model.Game, err error) {
	rdbGame := Game{
		GameName: g.Name,
	}
	result := repo.db.Connection.Create(&rdbGame)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to save: %s \n", err)
	}
	return repo.Find(rdbGame.ID)
}
