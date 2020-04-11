package controllers

import (
	"api-school/models"
	u "api-school/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateStudent = func(w http.ResponseWriter, r *http.Request) {

	student := &models.Student{}

	err := json.NewDecoder(r.Body).Decode(student)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := student.Create()
	u.Respond(w, resp)
}

var GetAllStudents = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetStudents()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetStudentByDocument = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	docNum, err := strconv.Atoi(params["doc_num"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetStudentByDocument(uint(docNum))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
