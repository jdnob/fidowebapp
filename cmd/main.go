package main

import (
	"context"
	"fidowebapp/config"
	"fidowebapp/database"
	"fmt"
	"html/template"
	"net/http"
	"os"

	dotenv "github.com/joho/godotenv"
)

func init() {
	_ = dotenv.Load()
}

type server struct{}

// api de base call using /
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))

	// case "POST":
	// 	w.WriteHeader(http.StatusCreated)
	// 	w.Write([]byte(`{"message": "post created"}`))
	// case "PUT":
	// 	w.WriteHeader(http.StatusAccepted)
	// 	w.Write([]byte(`{"message": "put created"}`))
	// case "DELETE":
	// 	w.WriteHeader(http.StatusAccepted)
	// 	w.Write([]byte(`{"message": "delete created"}`))
	// default:
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte(`{"message": "not found"}`))
	// }
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Status fine"}`))
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("static/index.html"))
	t.Execute(w, nil)
}

type PageVariables struct {
	Date string
	Time string
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/index.html")
}

func main() {
	// http.ListenAndServe(":9999", nil)

	ctx := context.Background()

	ctx = database.ContextWithDatabase(ctx, loadDatabaseConfig())

	db := database.DatabaseFromContext(ctx)

	user := database.FindUser(ctx, db)

	fmt.Printf("%v\n", user)
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
