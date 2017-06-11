//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: util.go
//     description:
//         created: 2017-06-10 14:55:17
//          author: wystan
//
//===============================================================================

package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/oswystan/bitmain/utils"
)

var logger = utils.Logger()

func sendResponse(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if v == nil {
		w.WriteHeader(code)
		return
	}

	bs, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bs = append(bs, '\n')
	w.WriteHeader(code)
	w.Write(bs)
}

func decodeBody(r *http.Request, v ...interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 8192))
	if err != nil {
		return err
	}
	r.Body.Close()

	for i := 0; i < len(v); i++ {
		if err = json.Unmarshal(body, v[i]); err != nil {
			return err
		}
	}

	return nil
}

//==================================== END ======================================
