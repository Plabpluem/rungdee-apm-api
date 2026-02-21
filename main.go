package main

import (
	"log"
	"os"
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/routes"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}
	dsn := os.Getenv("DATABASE_URL")
	config, err := pgx.ParseConfig(dsn)

	if err != nil {
		log.Fatalf("unable not parse %v", err)
	}

	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	dbs := stdlib.OpenDB(*config)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: dbs}), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("Can't connect database")
	}

	db.AutoMigrate(&entities.User{}, &entities.Customer{}, &entities.Contract{}, &entities.Room{}, &entities.Invoice{})

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	routes.UserRoutes(api, db)
	routes.RoomRoutes(api, db)

	app.Listen(":8080")
}
