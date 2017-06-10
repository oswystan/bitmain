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
)

func PostClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("do post class")
}

func GetTopTeacher(w http.ResponseWriter, r *http.Request) {
	log.Printf("do get top teacher")
}

//==================================== END ======================================
