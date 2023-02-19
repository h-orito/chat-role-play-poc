package application

import model "chat-role-play-poc/src/domain/model/game"

type GameService interface {
	Find(ID uint32) (game *model.Game, err error)
	Save(game model.Game) (saved *model.Game, err error)
}

type gameService struct {
	gameRepository model.GameRepository
}

func NewGameService(gameRepository model.GameRepository) GameService {
	return &gameService{
		gameRepository: gameRepository,
	}
}

func (gameService *gameService) Find(ID uint32) (game *model.Game, err error) {
	return gameService.gameRepository.Find(ID)
}

func (gameService *gameService) Save(game model.Game) (saved *model.Game, err error) {
	return gameService.gameRepository.Save(&game)
}
