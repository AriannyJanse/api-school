package main

import (
	"log"
	"net/http"
	"os"

	"api-test/app"
	"api-test/controllers"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "The element was not found"}`))
}

func main() {
	r := mux.NewRouter()
	r.Use(app.JwtAuthentication) //attach JWT auth middleware
	port := os.Getenv("app_port")
	if port == "" {
		port = "8000" //localhost
	}

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/user/login", controllers.Authenticate).Methods(http.MethodPost)
	api.HandleFunc("/user/new", controllers.CreateAccount).Methods(http.MethodPost)
	api.HandleFunc("/user/students", controllers.GetAllStudents).Methods(http.MethodGet)
	api.HandleFunc("/user/students/{doc_num}", controllers.GetStudentByDocument).Methods(http.MethodGet)
	api.HandleFunc("/user/students/new", controllers.CreateStudent).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
