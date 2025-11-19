package routes

import (
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type RouteConfig struct {
	App *fiber.App
}

func (c *RouteConfig) Setup() {
	v1 := c.App.Group("/api/v1")

	c.App.Use(middleware.CORSMiddleware())
	c.App.Use(recover.New())
	c.App.Use(healthcheck.New())

	// monitor
	v1.Get("/monitor", monitor.New())

}
