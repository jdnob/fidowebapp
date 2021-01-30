package api

import (
	"fidowebapp/handler"

	"github.com/gorilla/mux"
)

// TODO: To add later
func webAuthnRouter() *mux.Router {
	r := mux.NewRouter()
	// r.HandleFunc("/auth/make", BeginRegistration).Methods("GET")
	// r.HandleFunc("/auth/make", FinishRegistration).Methods("PUT")
	// r.HandleFunc("/auth/get", BeginLogin).Methods("GET")
	// r.HandleFunc("/auth/get", FinishLogin).Methods("GET")
	return r
}

func MyViewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.IndexHandler).Methods("GET")
	r.HandleFunc("/hello", handler.HelloWorldHandler).Methods("GET")
	r.HandleFunc("/error", handler.ErrorHandler).Methods("GET")
	return r
}
