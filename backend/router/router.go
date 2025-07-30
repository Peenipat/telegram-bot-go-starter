package router

import (
    "github.com/gofiber/fiber/v2"

    controller "github.com/Peenipat/telegram-bot-go-starter/backend/controller"
)

func RegisterTelegramRoutes(router fiber.Router, ctrl *controller.TelegramController) {
    
    telegramGroup := router.Group("/telegram")
	telegramGroup.Post("/webhook", ctrl.HandleWebhook)
	telegramGroup.Post("/send", ctrl.HandleSendMessage)

}