//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: router.go
//     description:
//         created: 2017-06-10 13:14:45
//          author: wystan
//
//===============================================================================

package router

import (
	"github.com/gorilla/mux"
	"github.com/oswystan/bitmain/controller"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/register-student", controller.PostStudent).Methods("POST")
	r.HandleFunc("/register-class", controller.PostClass).Methods("POST")
	r.HandleFunc("/get-class-total-score/{student_id:[0-9]+}", controller.GetSumScore).Methods("GET")
	r.HandleFunc("/get-top-teacher", controller.GetTopTeacher).Methods("GET")

	return r
}

//==================================== END ======================================
