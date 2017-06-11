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
	"net/http"

	"github.com/oswystan/bitmain/model"
	"github.com/oswystan/bitmain/store"
)

func PostClass(w http.ResponseWriter, r *http.Request) {
	class := &model.Class{}
	err := decodeBody(r, class)
	if err != nil {
		logger.Error(err)
		sendResponse(w, &model.Error{err.Error()}, http.StatusBadRequest)
		return
	}
	st := store.NewClassStore()
	err = st.Post(class)
	if err != nil {
		logger.Error(err)
		sendResponse(w, &model.Error{err.Error()}, http.StatusInternalServerError)
		return
	}
	sendResponse(w, class, http.StatusOK)
}

func GetTopTeacher(w http.ResponseWriter, r *http.Request) {
	st := store.NewClassStore()
	ret, err := st.GetTopTeacher()
	if err != nil {
		logger.Error(err)
		sendResponse(w, &model.Error{err.Error()}, http.StatusInternalServerError)
		return
	}
	sendResponse(w, ret, http.StatusOK)
}

//==================================== END ======================================
