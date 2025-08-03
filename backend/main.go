// @title Telegram Bot API
// @version 1.0
// @description This is a Telegram Bot API server
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"log"

	configDB "github.com/Peenipat/telegram-bot-go-starter/config"
	controller "github.com/Peenipat/telegram-bot-go-starter/controller"
	model  "github.com/Peenipat/telegram-bot-go-starter/model"
	swagger "github.com/swaggo/fiber-swagger"        
	_ "github.com/Peenipat/telegram-bot-go-starter/docs" 
	"github.com/Peenipat/telegram-bot-go-starter/router"
	"github.com/Peenipat/telegram-bot-go-starter/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)
func main() {
	configDB.LoadConfig()
	

	configDB.ConnectDB()
	if configDB.DB == nil {
		log.Fatal("GORM DB is nil. Cannot proceed.")
	}

	configDB.RunMigrations()
	configDB.SeedData()

	app := fiber.New()

	if err := configDB.DB.AutoMigrate(
		&model.MenuItem{},
		&model.Order{},
		&model.OrderItem{},
		&model.Reservation{},
		&model.Feedback{},
		&model.Ingredient{},
		&model.StockMovement{},
	); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://localhost:5174",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(compress.New())

	telegramService := service.NewTelegramService()
	telegramController := controller.NewTelegramController(telegramService)
	apiGroup := app.Group("/api/v1/")
	router.RegisterTelegramRoutes(apiGroup, telegramController)

	app.Get("/swagger/*", swagger.WrapHandler)

	port := configDB.AppConfig.Port
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
