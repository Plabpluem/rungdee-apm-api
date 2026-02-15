package routes

import (
	userAdapter "rungdee-apm-api/internal/adapters/user"
	userUsecases "rungdee-apm-api/internal/usecases/user"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func UserRoutes(app fiber.Router, db *gorm.DB) {
	userDb := userAdapter.NewGormUserRepository(db)
	userService := userUsecases.NewUserService(userDb)
	userHttp := userAdapter.NewHttpUserRepository(userService)

	app.Post("/user/signup", userHttp.Signup)
	app.Post("/user/login", userHttp.Login)
}
