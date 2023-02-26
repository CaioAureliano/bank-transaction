package handler

import (
	"encoding/json"
	"log"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/gofiber/fiber/v2"
)

type service interface {
	Create(dto.CreateRequestDTO) error
}

type Handler struct {
	s service
}

func New(s service) Handler {
	return Handler{s}
}

func (h Handler) CreateUser(c *fiber.Ctx) error {

	req := new(dto.CreateRequestDTO)

	if err := c.BodyParser(&req); err != nil {
		log.Printf("error to try parse request body - %s", err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}

	if errors := api.ValidateRequest(*req); errors != nil {
		errorsJson, _ := json.Marshal(errors)
		log.Printf("errors to try validate request body - %s", errorsJson)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := h.s.Create(*req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(201)
}
