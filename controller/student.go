//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: student.go
//     description:
//         created: 2017-06-10 13:24:38
//          author: wystan
//
//===============================================================================
package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswystan/bitmain/model"
	"github.com/oswystan/bitmain/store"
)

func PostStudent(w http.ResponseWriter, r *http.Request) {
	student := &model.Student{}
	err := decodeBody(r, student)
	if err != nil {
		logger.Error(err)
		sendResponse(w, &model.Error{err.Error()}, http.StatusBadRequest)
		return
	}

	st := store.NewStudentStore()
	err = st.Post(student)
	if err != nil {
		logger.Error(err)
		sendResponse(w, &model.Error{err.Error()}, http.StatusInternalServerError)
		return
	}
	sendResponse(w, student, http.StatusOK)
}

func GetSumScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, ok := vars["student_id"]
	if !ok {
		sendResponse(w, &model.Error{"no student id found"}, http.StatusBadRequest)
		return
	}
	student := &model.Student{Id: sid}
	st := store.NewStudentStore()
	score, err := st.GetTotalScore(student)
	if err != nil {
		logger.Error(err)
		sendResponse(w, &model.Error{err.Error()}, http.StatusInternalServerError)
		return
	}

	sendResponse(w, score, http.StatusOK)
}

//==================================== END ======================================
