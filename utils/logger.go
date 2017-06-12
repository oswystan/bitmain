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
	logger.AddFilter("file", l4g.FINE, flw)
	logger.Info("################################")
	logger.Info("starting server ....            ")
	logger.Info("################################")
}

func LogExit() {
	logger.Warn("################################")
	logger.Warn("server exit !")
	logger.Warn("################################")
	logger.Close()
	time.Sleep(100 * time.Millisecond)
}

//==================================== END ======================================
