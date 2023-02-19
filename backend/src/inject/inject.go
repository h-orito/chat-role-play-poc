package inject

import (
	"chat-role-play-poc/src/adapter/controller"
	application "chat-role-play-poc/src/application/service"
	model "chat-role-play-poc/src/domain/model/game"
	db "chat-role-play-poc/src/infrastructure/rdb"
)

// routing
func InjectRouting() *controller.Routing {
	gameController := injectGameController()
	return controller.NewRouting(gameController)
}

// controller
func injectGameController() controller.GameController {
	gameService := injectGameService()
	return controller.NewGameController(gameService)
}

// service
func injectGameService() application.GameService {
	gameRepository := injectGameRepository()
	return application.NewGameService(gameRepository)
}

// repository
func injectGameRepository() model.GameRepository {
	database := injectDb()
	return db.NewGameRepository(&database)
}

// database
func injectDb() db.DB {
	return db.NewDB()
}
