package controllers

import (
	"api-school/models"
	u "api-school/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateTeacher = func(w http.ResponseWriter, r *http.Request) {

	teacher := &models.Teacher{}

	err := json.NewDecoder(r.Body).Decode(teacher)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := teacher.Create()
	u.Respond(w, resp)
}

var GetAllTeachers = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetTeachers()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetTeacherByDocument = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	docNum, err := strconv.Atoi(params["doc_num"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetTeacherByDocument(uint(docNum))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteTeacherByDocNum = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	docNum, err := strconv.Atoi(params["doc_num"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	u.Respond(w, models.DeleteTeacherByDocNum(uint(docNum)))
}
