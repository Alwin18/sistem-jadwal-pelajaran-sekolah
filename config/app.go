package config

import (
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/handlers"
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/routes"
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Cfg      *Config
}

func Bootstrap(cfg *BootstrapConfig) {
	// initiate service
	authService := services.NewAuthService(cfg.DB, cfg.Log)
	scheduleService := services.NewScheduleService(cfg.DB, cfg.Log)

	// initiate handlers
	authHandler := handlers.NewAuthHandler(authService, cfg.Validate)
	scheduleHandler := handlers.NewScheduleHandler(scheduleService, cfg.Validate)

	routeConfig := routes.RouteConfig{
		App:      cfg.App,
		Auth:     authHandler,
		Schedule: scheduleHandler,
	}
	routeConfig.Setup()
}
