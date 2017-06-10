//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: database.go
//     description:
//         created: 2017-06-10 12:22:29
//          author: wystan
//
//===============================================================================
package store

import (
	"fmt"

	"github.com/go-pg/pg"
)

var db *pg.DB = nil

func Connect(db_name, user, passwd string) error {
	if db != nil {
		return nil
	}
	opt := pg.Options{
		User:     user,
		Password: passwd,
		Database: db_name,
	}

	db = pg.Connect(&opt)
	if db == nil {
		return fmt.Errorf("fail to connect database %s", db_name)
	}
	return nil
}

func Close() {
	if db != nil {
		db.Close()
		db = nil
	}
}

//==================================== END ======================================
