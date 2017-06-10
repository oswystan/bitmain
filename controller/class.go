//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: class.go
//     description:
//         created: 2017-06-10 13:48:58
//          author: wystan
//
//===============================================================================

package controller

import (
	"log"
	"net/http"

	"github.com/oswystan/bitmain/model"
)

func PostClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("do post class")
	class := &model.Class{}
	err := decodeBody(r, class)
	if err != nil {
		sendResponse(w, &model.Error{err.Error()}, http.StatusBadRequest)
		return
	}
	sendResponse(w, class, http.StatusOK)
}

func GetTopTeacher(w http.ResponseWriter, r *http.Request) {
	log.Printf("do get top teacher")
	teacher := &model.TopTeacher{"lily"}
	sendResponse(w, teacher, http.StatusOK)
}

//==================================== END ======================================
