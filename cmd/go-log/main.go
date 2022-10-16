package main

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/pkg/errors"
)

func main() {

	log.DebugMode = true
	log.PrintTimestamp = true
	log.PrintColors = true
	log.TimeFormat = "2006-01-02 15:04:05.000"

	myVar := map[string]string{"hello": "world"}

	log.Debug("debug arg1", "debug arg2")
	log.Debugf("debug arg1 %d", 1)
	log.DebugDump(myVar, "debug prefix")
	log.DebugSeparator("debug title")

	log.Info("info arg1", "info arg2")
	log.Infof("info arg1 %d", 1)
	log.InfoDump(myVar, "info prefix")
	log.InfoSeparator("info title")

	log.Warn("warn arg1", "warn arg2")
	log.Warnf("warn arg1 %d", 1)
	log.WarnDump(myVar, "warn prefix")
	log.WarnSeparator("warn title")

	log.Error("error arg1", "error arg2")
	log.Errorf("error arg1 %d", 1)
	log.ErrorDump(myVar, "error prefix")
	log.ErrorSeparator("error title")
	log.Error(errors.New("error"))

	log.StackTrace(errors.New("error with stack trace"))

	log.Fatal("fatal arg1", "fatal arg2")
	log.Fatalf("fatal arg1 %d", 1)

}
