package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/status", getStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
