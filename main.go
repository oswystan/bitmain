//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: main.go
//     description:
//         created: 2017-06-10 12:21:09
//          author: wystan
//
//===============================================================================

package main

import (
	"net/http"

	"github.com/oswystan/bitmain/router"
	"github.com/oswystan/bitmain/store"
	"github.com/oswystan/bitmain/utils"
)

func main() {
	logger := utils.Logger()
	logger.Info("starting server...")
	err := store.Connect("bitmain", "bitmain", "bitmain")
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer store.Close()
	logger.Info("database connected")

	r := router.NewRouter()
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		logger.Error("ERROR: %s", err.Error())
		return
	}

	logger.Warn("Server Exit!")
}

//==================================== END ======================================
