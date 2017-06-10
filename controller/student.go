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
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswystan/bitmain/model"
)

func PostStudent(w http.ResponseWriter, r *http.Request) {
	log.Printf("post student")
	student := &model.Student{}
	err := decodeBody(r, student)
	if err != nil {
		sendResponse(w, &model.Error{err.Error()}, http.StatusBadRequest)
		return
	}
	st := &model.Student{
		Id: "1000",
	}

	sendResponse(w, st, http.StatusOK)
}

func GetSumScore(w http.ResponseWriter, r *http.Request) {
	log.Printf("get score")
	vars := mux.Vars(r)
	_, ok := vars["student_id"]
	if !ok {
		sendResponse(w, &model.Error{"no student id found"}, http.StatusBadRequest)
		return
	}

	score := &model.TotalScore{100}
	sendResponse(w, score, http.StatusOK)
}

//==================================== END ======================================
