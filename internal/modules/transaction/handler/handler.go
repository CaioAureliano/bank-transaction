package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/gofiber/fiber/v2"
)

type service interface {
	CreateTransaction(*dto.TransactionRequestDTO) (uint, error)
}

type Handler struct {
	s service
}

func New(s service) Handler {
	return Handler{s}
}

func (h Handler) CreateTransaction(c *fiber.Ctx) error {

	req := new(dto.TransactionRequestDTO)

	if err := c.BodyParser(&req); err != nil {
		log.Printf("error to try parse request body - %s", err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}

	if errors := api.ValidateRequest(*req); errors != nil {
		errorsJson, _ := json.Marshal(errors)
		log.Printf("errors to try validate request body - %s", errorsJson)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	id, err := h.s.CreateTransaction(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Transaction Requested",
		"links": fiber.Map{
			"href": fmt.Sprintf("/trasactions/%d", id),
			"rel":  "transactions",
			"type": "GET",
		},
	})
}
