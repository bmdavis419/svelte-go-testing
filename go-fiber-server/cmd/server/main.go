package main

import (
	"context"
	"fmt"

	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/private/db"
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/private/handlers"
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/private/storage"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func newFiberServer(lc fx.Lifecycle, userHandlers *handlers.UserHandler) *fiber.App {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// attach the user handlers
	userGroup := app.Group("/users")
	userGroup.Post("/generate", userHandlers.GenerateNewUser)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// TODO: switch the port to an env variable
			fmt.Println("Starting fiber server on port 8080")
			go app.Listen(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}

func main() {
	fx.New(
		fx.Provide(db.CreatePostgresConnection, storage.NewUserStorage, handlers.NewUserHandler),
		fx.Invoke(newFiberServer),
	).Run()
}
