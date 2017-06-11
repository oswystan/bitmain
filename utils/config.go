//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: config.go
//     description:
//         created: 2017-06-11 23:23:35
//          author: wystan
//
//===============================================================================
package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Database string `json:"database"`
	User     string `json:"user"`
	Passwd   string `json:"password"`
}

func LoadConfig(out *Config) error {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bytes, out); err != nil {
		return err
	}
	return nil
}

//==================================== END ======================================
