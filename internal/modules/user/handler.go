package user

import (
	"github.com/CaioAureliano/bank-transaction/internal/shared/api"
	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		req := new(CreateRequestDTO)

		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
		}

		if errors := api.ValidateRequest(*req); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}

		return c.SendStatus(201)
	}
}
