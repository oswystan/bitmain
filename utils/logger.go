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
	"time"

	l4g "github.com/alecthomas/log4go"
)

var logger = l4g.NewDefaultLogger(l4g.FINE)

func Logger() l4g.Logger {
	return logger
}

func init() {
	flw := l4g.NewFileLogWriter("bitmain.log", false)
	flw.SetRotate(true)
	flw.SetFormat("[%D %T] [%L]%M")
	flw.SetRotateSize(2 * 1024 * 1024)
	flw = flw.SetRotateMaxBackup(10)
	flw.SetHeadFoot("---------------- START --------------", "---------------- END ----------------")
	logger.AddFilter("file", l4g.FINE, flw)
}

func LogExit() {
	logger.Warn("server exit !")
	logger.Close()
	time.Sleep(100 * time.Millisecond)
}

//==================================== END ======================================
