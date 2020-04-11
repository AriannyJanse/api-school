package models

import (
	u "api-school/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

//A struct to rep students
type Teacher struct {
	gorm.Model
	Name    string `json:"name"`
	DocType string `json:"doc_type"`
	DocNum  uint   `json:"doc_num"`
	Course  string `json:"course"`
	UserId  uint   `json:"user_id"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (teacher *Teacher) Validate() (map[string]interface{}, bool) {

	if teacher.Name == "" {
		return u.Message(false, "Student name should be on the payload"), false
	}

	if teacher.DocType == "" {
		return u.Message(false, "Document type should be on the payload"), false
	}

	if teacher.DocNum <= 0 {
		return u.Message(false, "Document number shuold be on the payload"), false
	}

	if teacher.Course <= "" {
		return u.Message(false, "Course shuold be on the payload"), false
	}

	if teacher.UserId <= 0 {
		return u.Message(false, "The user was not recognize"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (teacher *Teacher) Create() map[string]interface{} {

	if resp, ok := teacher.Validate(); !ok {
		return resp
	}

	GetDB().Create(teacher)

	resp := u.Message(true, "success")
	resp["teacher"] = teacher
	return resp
}

func GetTeacherByDocument(documentNum uint) *Teacher {

	teacher := &Teacher{}
	err := GetDB().Table("teachers").Where("doc_num = ?", documentNum).First(teacher).Error
	if err != nil {
		return nil
	}
	return teacher
}

func GetTeachers() []*Teacher {

	teachers := make([]*Teacher, 0)
	err := GetDB().Table("teachers").Find(&teachers).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return teachers
}
