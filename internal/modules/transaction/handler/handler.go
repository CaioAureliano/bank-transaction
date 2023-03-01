package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/CaioAureliano/bank-transaction/pkg/errors"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type service interface {
	CreateTransaction(req *dto.TransactionRequestDTO, userID uint) (uint, error)
	GetTransaction(*dto.GetTransactionRequestDTO) *dto.TransactionResponseDTO
}

type Handler struct {
	s service
}

func New(s service) Handler {
	return Handler{s}
}

// @Summary 	Create Transaction
// @Description receive payload to do transfer if is valid send message to queue
// @Accept  	json
// @Produce  	json
// @Tags 		transactions
// @Param		transaction body	  dto.TransactionRequestDTO			true	"transaction data"
// @Success 	202			{object}  dto.CreatedTransactionResponseDTO
// @Failure 	400 		{object}  errors.HttpErrorResponse
// @Failure 	401 		{object}  errors.HttpErrorResponse
// @Failure 	422 		{object}  errors.HttpErrorResponse
// @Failure 	500 		{object}  errors.HttpErrorResponse
// @Security 	JwtToken
// @Router 		/transactions  [post]
func (h Handler) CreateTransaction(c *fiber.Ctx) error {

	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	typeAccount := claims["type"].(float64)
	userID := claims["ID"].(float64)

	req := new(dto.TransactionRequestDTO)

	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors.NewHttpError("failed to parse request body", fiber.StatusUnprocessableEntity, err.Error()))
	}

	if err := api.ValidateRequest(*req); err != nil {
		errorsJson, _ := json.Marshal(err)
		log.Println(string(errorsJson))
		return c.Status(fiber.StatusBadRequest).JSON(errors.NewHttpError("failed to validate request", fiber.StatusBadRequest, string(errorsJson)))
	}

	if model.Type(typeAccount) != model.USER {
		return c.Status(fiber.StatusUnauthorized).JSON(errors.NewHttpError("invalid payer", fiber.StatusUnauthorized, "user not have permission to do transaction"))
	}

	if req.Payee == uint(userID) || req.Value <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errors.NewHttpError("bad request", fiber.StatusBadRequest, "nice try"))
	}

	id, err := h.s.CreateTransaction(req, uint(userID))
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(errors.NewHttpError("failed to create transaction", fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusAccepted).JSON(dto.CreatedTransactionResponseDTO{
		Message: "transaction request",
		Links: dto.LinksHateoas{
			Href: fmt.Sprintf("/trasactions/%d", id),
			Rel:  "transactions",
			Type: "GET",
		},
	})
}

// @Summary 	Get Transaction
// @Description get cached transaction status to short polling(without error return, just empty string or status)
// @Accept  	json
// @Produce  	json
// @Tags		transactions
// @Param		id  path	  int				 		 	true	"transaction id"
// @Success 	200 {object}  dto.TransactionResponseDTO
// @Security 	JwtToken
// @Router 		/transactions/:id  [get]
func (h Handler) GetTransaction(c *fiber.Ctx) error {

	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	userID := claims["ID"].(float64)

	transactionID := c.Params("id")

	req := &dto.GetTransactionRequestDTO{
		TransactionID: transactionID,
		PayerID:       uint(userID),
	}

	res := h.s.GetTransaction(req)

	return c.JSON(res)
}
