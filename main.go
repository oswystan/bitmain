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
	defer utils.LogExit()
	logger.Info("starting server...")
	cfg := &utils.Config{}
	err := utils.LoadConfig(cfg)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	err = store.Connect(cfg.Database, cfg.User, cfg.Passwd)
	if err != nil {
		logger.Error("fail to connect database: %s", err.Error())
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
}

//==================================== END ======================================
