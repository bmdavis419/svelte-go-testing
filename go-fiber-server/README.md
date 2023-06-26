## what this project uses

1. MySQL, the connection string is set in the `Taskfile.yaml` file, make sure to run the migrations before running the project.
2. Redis, the default connection string is `localhost:6379`, you can change it in the `db/redis.go` file.

## todo

1. Refactor the auth check into either a middleware or more concise function.
2. Improve the error handling.
3. Convert all config to use a `.env` file.