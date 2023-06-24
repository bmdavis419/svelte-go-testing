package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func CreatePostgresConnection(lc fx.Lifecycle) *pgxpool.Pool {
	conn, err := pgxpool.New(context.TODO(), os.Getenv("POSTGRESQL_URL"))
	fmt.Println("Connected to postgres", conn, os.Getenv("POSTGRESQL_URL"))

	// 50/50 shot this is goofy, will consult the viewers smarter than me :)
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			if err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			conn.Close()
			return nil
		},
	})

	return conn
}
