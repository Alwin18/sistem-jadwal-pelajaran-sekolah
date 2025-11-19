package config

import (
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/routes"
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

	// initiate handlers

	routeConfig := routes.RouteConfig{
		App: cfg.App,
	}
	routeConfig.Setup()
}
