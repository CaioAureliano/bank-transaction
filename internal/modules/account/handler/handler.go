package handler

import (
	"encoding/json"
	"log"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/CaioAureliano/bank-transaction/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type service interface {
	CreateUserAccount(dto.CreateRequestDTO) error
	Authenticate(dto.AuthRequestDTO) (string, error)
}

type Handler struct {
	s service
}

func New(s service) Handler {
	return Handler{s}
}

// @Summary 	Create User Account
// @Description receive payload with user data and if valid create user account
// @Accept  	json
// @Produce  	json
// @Tags 		accounts
// @Param		user body	   dto.CreateRequestDTO	true	"user data"
// @Success 	201  {integer} int
// @Failure 	400  {object}  errors.HttpErrorResponse
// @Failure 	422  {object}  errors.HttpErrorResponse
// @Failure 	500  {object}  errors.HttpErrorResponse
// @Router 		/accounts [post]
func (h Handler) CreateUser(c *fiber.Ctx) error {

	req := new(dto.CreateRequestDTO)

	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors.NewHttpError("failed to parse request body", fiber.StatusUnprocessableEntity, err.Error()))
	}

	if err := api.ValidateRequest(*req); err != nil {
		errorsJson, _ := json.Marshal(err)
		log.Println(string(errorsJson))
		return c.Status(fiber.StatusBadRequest).JSON(errors.NewHttpError("invalid request", fiber.StatusBadRequest, string(errorsJson)))
	}

	if err := h.s.CreateUserAccount(*req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(errors.NewHttpError("failed to create user account", fiber.StatusInternalServerError, err.Error()))
	}

	return c.SendStatus(201)
}

// @Summary 	Authenticate User
// @Description Receive email with password and if is valid than return JWT
// @Accept  	json
// @Produce  	json
// @Tags 		accounts
// @Param		user body	  dto.AuthRequestDTO		true	"authentication payload"
// @Success 	200  {object} dto.JwtResponseDTO
// @Failure 	400  {object} errors.HttpErrorResponse
// @Failure 	401  {object} errors.HttpErrorResponse
// @Failure 	422  {object} errors.HttpErrorResponse
// @Router 		/accounts/auth [post]
func (h Handler) Authenticate(c *fiber.Ctx) error {

	req := new(dto.AuthRequestDTO)

	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors.NewHttpError("failed to parse request body", fiber.StatusUnprocessableEntity, err.Error()))
	}

	if err := api.ValidateRequest(*req); err != nil {
		errorsJson, _ := json.Marshal(err)
		log.Println(string(errorsJson))
		return c.Status(fiber.StatusBadRequest).JSON(errors.NewHttpError("invalid request", fiber.StatusBadRequest, string(errorsJson)))
	}

	token, err := h.s.Authenticate(*req)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(errors.NewHttpError("error to try authenticate", fiber.StatusUnauthorized, err.Error()))
	}

	return c.JSON(dto.JwtResponseDTO{Token: token})
}
