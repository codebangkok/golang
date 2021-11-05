package handlers

import "github.com/gofiber/fiber/v2"

type CatalogHandler interface {
	GetProducts(c *fiber.Ctx) error
}
