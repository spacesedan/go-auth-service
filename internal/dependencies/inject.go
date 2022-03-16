package dependencies

import (
	"authentication/internal/handlers"
	"authentication/internal/repo"
	"authentication/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Inject register app services with database connection
func Inject(db *mongo.Database) (*gin.Engine, error) {
	// init dao
	dao := repo.NewDAO(db)

	// register services
	tokenService := service.NewTokenService()
	userService := service.NewUserService(dao)

	app := gin.Default()

	config := &handlers.Config{
		UserService:  userService,
		TokenService: tokenService,
		R:            app,
	}

	handlers.NewHandler(config)

	return app, nil
}
