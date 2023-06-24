package main

import (
	"context"
	"fmt"

	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/auth"
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

//	@title			Go Svelte Todos API
//	@version		1.0
//	@description	This is a basic example API.

// @host		localhost:8080
// @BasePath	/
func newFiberServer(lc fx.Lifecycle, userHandlers *handlers.UserHandler) *fiber.App {
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
		fx.Provide(
			// creates: *pgxpool.Pool
			db.CreatePostgresConnection,
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

// WAS TESTING OUT MYSQL
// func dbDemo(db *sqlx.DB) {
// 	mostRecentForm := make([]struct {
// 		Acc_num          string
// 		Period_of_report string
// 		Issuer_cik       string
// 	}, 0)
// 	err := db.Select(&mostRecentForm, "SELECT acc_num, period_of_report, issuer_cik FROM form ORDER BY period_of_report DESC LIMIT 10")
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, row := range mostRecentForm {
// 		fmt.Printf("Form: %s, on %s by %s\n", row.Acc_num, row.Period_of_report, row.Issuer_cik)
// 	}
// }
