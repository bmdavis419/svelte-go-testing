package handlers

import (
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/private/storage"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Storage *storage.UserStorage
}

func NewUserHandler(storage *storage.UserStorage) *UserHandler {
	return &UserHandler{Storage: storage}
}

func (u *UserHandler) GenerateNewUser(c *fiber.Ctx) error {
	id, err := u.Storage.GenerateNewUser()
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"id": id})
}
