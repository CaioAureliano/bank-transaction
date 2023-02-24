package handler

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/user/domain/dto"
	"github.com/CaioAureliano/bank-transaction/internal/modules/user/service"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/gofiber/fiber/v2"
)

var userService = service.NewService

func CreateUser(c *fiber.Ctx) error {

	req := new(dto.CreateRequestDTO)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}

	if errors := api.ValidateRequest(*req); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := userService().Create(*req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(201)
}
