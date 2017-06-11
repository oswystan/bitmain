//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: logger.go
//     description:
//         created: 2017-06-11 11:51:22
//          author: wystan
//
//===============================================================================

package utils

import (
	l4g "github.com/keepeye/log4go"
)

var logger = l4g.NewDefaultLogger(l4g.FINE)

func Logger() l4g.Logger {
	return logger
}

func init() {
	logger.AddFilter("file", l4g.FINE, l4g.NewFileLogWriter("bitmain.log", false))
}

//==================================== END ======================================
