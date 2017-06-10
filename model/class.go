//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: class.go
//     description:
//         created: 2017-06-10 15:29:30
//          author: wystan
//
//===============================================================================

package model

type Class struct {
	ClassNumber int    `json:"classNumber"`
	Teacher     string `json:"teacher"`
}

type TopTeacher struct {
	Teacher string `json:"teacher"`
}

//==================================== END ======================================
