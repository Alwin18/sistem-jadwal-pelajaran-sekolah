package handlers

import (
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/dto"
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/services"
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ScheduleHandler struct {
	ScheduleService *services.ScheduleService
	validate        *validator.Validate
}

func NewScheduleHandler(ScheduleService *services.ScheduleService, validate *validator.Validate) *ScheduleHandler {
	return &ScheduleHandler{
		ScheduleService: ScheduleService,
		validate:        validate,
	}
}

func (h *ScheduleHandler) ListTeacher(ctx *fiber.Ctx) error {
	var param dto.ListScheduleRequest
	if err := ctx.QueryParser(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewErrorResponse(utils.ResponseError{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		}))
	}

	if err := h.validate.Struct(param); err != nil {
		errors := utils.FormatValidationErrors(err, param)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewErrorResponse(utils.ResponseError{
			Message: "param request tidak sesuai",
			Code:    fiber.StatusBadRequest,
			Errors:  errors,
		}))
	}

	resp, err := h.ScheduleService.ListTeacher(ctx, param)
	if err != nil {
		return ctx.JSON(utils.NewErrorResponse(utils.ResponseError{
			Message: err.Error(),
			Code:    ctx.Response().Header.StatusCode(),
		}))
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.NewResponse(resp, "success get list schedule", fiber.StatusOK))
}
