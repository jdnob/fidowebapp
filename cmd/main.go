package main

import (
	"context"
	"fidowebapp/api"
	"fidowebapp/appcontext"
	"fidowebapp/config"
	"fidowebapp/database"
	"fidowebapp/entity"
	"flag"
	"net/http"

	"os"

	dotenv "github.com/joho/godotenv"
)

func init() {
	_ = dotenv.Load()
}

type Args struct {
	LogLevel string
}

func parseArgs() Args {
	var args Args
	flag.StringVar(&args.LogLevel, "log", "info", "Log level [trace, debug, info, warning, error]")
	flag.Parse()
	return args
}

func main() {
	args := parseArgs()

	ctx := context.Background()

	ctx = database.ContextWithDatabase(ctx, loadDatabaseConfig())

	db := database.DatabaseFromContext(ctx)
	appcontext.SetupLog(ctx, args.LogLevel, true)

	entity.FindAllUsers(ctx, db)
	entity.FindUser(ctx, db, "4bacf836-3d6d-401e-99dc-54879cab1975")

	r := api.MyViewRouter()
	http.ListenAndServe(":9999", r)
}

func loadDatabaseConfig() config.DatabaseConfiguration {
	return config.DatabaseConfiguration{
		DBName:     os.Getenv("FWA_DATABASE_NAME"),
		DBURL:      os.Getenv("FWA_DATABASE_HOST"),
		DBPassword: os.Getenv("FWA_DATABASE_PASSWORD"),
		DBPort:     os.Getenv("FWA_DATABASE_PORT"),
		DBUser:     os.Getenv("FWA_DATABASE_USER"),
	}
}
