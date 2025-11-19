package routes

import (
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/handlers"
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type RouteConfig struct {
	App      *fiber.App
	Auth     *handlers.AuthHandler
	Schedule *handlers.ScheduleHandler
}

func (c *RouteConfig) Setup() {
	v1 := c.App.Group("/api/v1")

	c.App.Use(middleware.CORSMiddleware())
	c.App.Use(recover.New())
	c.App.Use(healthcheck.New())

	// monitor
	v1.Get("/monitor", monitor.New())

	AuthRoute(c, v1.Group("auth"))
	ScheduleRoute(c, v1.Group("schedule"))
}

func AuthRoute(c *RouteConfig, v1 fiber.Router) {
	v1.Post("login", c.Auth.Login)
}

func ScheduleRoute(c *RouteConfig, v1 fiber.Router) {
	v1.Get("teacher", c.Schedule.ListTeacher)
}
