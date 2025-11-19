package handlers

import (
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/dto"
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/services"
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *services.AuthService
	validate    *validator.Validate
}

func NewAuthHandler(authService *services.AuthService, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validate,
	}
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	var loginRequest dto.LoginRequest
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewErrorResponse(utils.ResponseError{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		}))
	}

	if err := h.validate.Struct(loginRequest); err != nil {
		errors := utils.FormatValidationErrors(err, loginRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewErrorResponse(utils.ResponseError{
			Message: "body request tidak sesuai",
			Code:    fiber.StatusBadRequest,
			Errors:  errors,
		}))
	}

	resp, err := h.authService.Login(ctx, loginRequest.Username, loginRequest.Password)
	if err != nil {
		return ctx.JSON(utils.NewErrorResponse(utils.ResponseError{
			Message: err.Error(),
			Code:    ctx.Response().Header.StatusCode(),
		}))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponse(resp, "success login", fiber.StatusOK))
}
