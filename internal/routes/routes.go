package routes

import (
	roomAdapters "rungdee-apm-api/internal/adapters/room"
	userAdapter "rungdee-apm-api/internal/adapters/user"
	roomUsecases "rungdee-apm-api/internal/usecases/room"
	userUsecases "rungdee-apm-api/internal/usecases/user"
	"rungdee-apm-api/pkg/middleware"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func UserRoutes(app fiber.Router, db *gorm.DB) {
	userDb := userAdapter.NewGormUserRepository(db)
	userService := userUsecases.NewUserService(userDb)
	userHttp := userAdapter.NewHttpUserHandler(userService)

	app.Post("/user/login", userHttp.Login)
	app.Post("/user/signup", userHttp.Signup)
}

func RoomRoutes(app fiber.Router, db *gorm.DB) {
	roomDb := roomAdapters.NewGormRoomRepository(db)
	roomService := roomUsecases.NewRoomService(roomDb)
	roomHttp := roomAdapters.NewHttpRoomHandler(roomService)

	app.Use("/room", middleware.AuthRequired, middleware.RbacRequired([]string{"admin", "employee"}))

	app.Get("/room", roomHttp.FindAll)
	app.Get("/room/:id", roomHttp.Find)
	app.Post("/room", roomHttp.Create)
}
