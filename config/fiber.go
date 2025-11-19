package config

import "github.com/gofiber/fiber/v2"

func NewFiber(cfg *Config) *fiber.App {
	return fiber.New(fiber.Config{
		Network:        fiber.NetworkTCP,
		ReadBufferSize: 32 * 1024,        // 32 KB — naikkan dari default 4 KB
		BodyLimit:      25 * 1024 * 1024, // 25 MB (opsional, aman untuk upload)
	})
}
