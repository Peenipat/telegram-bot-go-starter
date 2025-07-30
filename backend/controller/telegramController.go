package controller

import (
	"github.com/gofiber/fiber/v2"
	Interface "github.com/Peenipat/telegram-bot-go-starter/backend/interface"
)

type TelegramController struct {
	Service Interface.ITelegramService
}

func NewTelegramController(service Interface.ITelegramService) *TelegramController {
	return &TelegramController{Service: service}
}

func (ctl *TelegramController) HandleWebhook(c *fiber.Ctx) error {
	return ctl.Service.ProcessWebhook(c)
}


func (ctl *TelegramController) HandleSendMessage(c *fiber.Ctx) error {
	var req Interface.TelegramSendRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid payload",
		})
	}

	if req.ChatID == 0 || req.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "chat_id and message are required",
		})
	}

	err := ctl.Service.SendTelegramMessage(req.ChatID, req.Message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "sent",
	})
}
