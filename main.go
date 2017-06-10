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
	"log"
	"net/http"

	"github.com/oswystan/bitmain/router"
	"github.com/oswystan/bitmain/store"
)

func main() {
	log.Printf("starting server...")
	err := store.Connect("bitmain", "bitmain", "bitmain")
	if err != nil {
		log.Printf(err.Error())
	}
	defer store.Close()
	log.Printf("database connected")

	r := router.NewRouter()
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return
	}

	log.Printf("Server Exit!")
}

//==================================== END ======================================
