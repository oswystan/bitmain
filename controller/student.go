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
)

func PostStudent(w http.ResponseWriter, r *http.Request) {
	log.Printf("post student")
}

func GetSumScore(w http.ResponseWriter, r *http.Request) {
	log.Printf("get score")
}

//==================================== END ======================================
