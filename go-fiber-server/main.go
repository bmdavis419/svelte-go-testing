package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/auth"
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/config"
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/db"
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/handlers"
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"

	_ "github.com/bmdavis419/svelte-go-testing/go-fiber-server/docs"
	_ "github.com/go-sql-driver/mysql"
)

//		@title			Go Svelte Todos API
//		@version		1.0
//		@description	This is a basic CRUD api with authentication, written by @bmdavis419.
//		@securityDefinitions.apikey	ApiKeyAuth
//		@in							header
//		@name						Authorization
//		@description				Token in Bearer format to authenticate the user
//		@host		localhost:8080
//	 @BasePath	/
func newFiberServer(lc fx.Lifecycle, userHandlers *handlers.UserHandler, todoHandlers *handlers.TodoHandler) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// docs
	app.Get("/swagger/*", swagger.HandlerDefault)

	// attach the user handlers
	userGroup := app.Group("/users")
	userGroup.Post(("/sign-up"), userHandlers.SignUpUser)
	userGroup.Post("/sign-in", userHandlers.SignInUser)
	userGroup.Get("/me", userHandlers.GetUserInfo)
	userGroup.Post("/sign-out", userHandlers.SignOutUser)

	// attach the todo handlers
	todoGroup := app.Group("/todos")
	todoGroup.Post("/", todoHandlers.CreateTodo)
	todoGroup.Get("/", todoHandlers.FetchTodos)
	todoGroup.Get("/:id", todoHandlers.FetchTodo)
	todoGroup.Put("/:id/complete", todoHandlers.CompleteTodo)
	todoGroup.Delete("/:id", todoHandlers.DeleteTodo)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			port := os.Getenv("PORT")
			if port == "" {
				port = "8080"
			}
			fmt.Printf("Starting fiber server on port %s\n", port)
			go app.Listen(":" + port)
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
		fx.Provide(
			// creates: config.EnvVars
			config.LoadEnv,
			// creates: *storage.TodoStorage
			storage.NewTodoStorage,
			// creates: *handlers.TodoHandler
			handlers.NewTodoHandler,
			// creates: *sqlx.DB
			db.CreateMySqlConnection,
			// creates: *storage.UserStorage
			storage.NewUserStorage,
			// creates: *handlers.UserHandler
			handlers.NewUserHandler,
			// creates: *redis.Client
			db.CreateRedisConnection,
			// creates: *auth.SessionManager
			auth.NewSessionManager),
		fx.Invoke(newFiberServer),
	).Run()
}
