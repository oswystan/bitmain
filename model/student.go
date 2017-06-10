//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: student.go
//     description:
//         created: 2017-06-10 14:59:18
//          author: wystan
//
//===============================================================================

package model

type Student struct {
	Id          string `json:"id"`
	ClassNumber int    `json:"classNumber"`
	Score       int    `json:"score"`
}

type TotalScore struct {
	Total int `json:"total"`
}

//==================================== END ======================================
