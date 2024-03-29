package main

import (
	"log"
	"net/http"
	"os"

	"api-school/app"
	"api-school/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "The element was not found"}`))
}

func main() {
	r := mux.NewRouter()
	r.Use(app.JwtAuthentication) //attach JWT auth middleware
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/user/login", controllers.Authenticate).Methods(http.MethodPost)
	api.HandleFunc("/user/new", controllers.CreateAccount).Methods(http.MethodPost)

	api.HandleFunc("/user/students", controllers.GetAllStudents).Methods(http.MethodGet)
	api.HandleFunc("/user/students/{doc_num}", controllers.GetStudentByDocument).Methods(http.MethodGet)
	api.HandleFunc("/user/students/new", controllers.CreateStudent).Methods(http.MethodPost)
	api.HandleFunc("/user/students/{doc_num}", controllers.DeleteStudentByDocNum).Methods(http.MethodDelete)
	api.HandleFunc("/user/students/{doc_num}", controllers.UpdateStudentByDocNum).Methods(http.MethodPut)

	api.HandleFunc("/user/teachers", controllers.GetAllTeachers).Methods(http.MethodGet)
	api.HandleFunc("/user/teachers/{doc_num}", controllers.GetTeacherByDocument).Methods(http.MethodGet)
	api.HandleFunc("/user/teachers/new", controllers.CreateTeacher).Methods(http.MethodPost)
	api.HandleFunc("/user/teachers/{doc_num}", controllers.DeleteTeacherByDocNum).Methods(http.MethodDelete)
	api.HandleFunc("/user/teachers/{doc_num}", controllers.UpdateTeacherByDocNum).Methods(http.MethodPut)

	//cors optionsGoes Below
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("allowed_origins")}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application

		},
	})

	log.Fatal(http.ListenAndServe(":"+port, corsOpts.Handler(r)))
}
