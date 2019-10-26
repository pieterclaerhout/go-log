package main

import (
	"github.com/pieterclaerhout/go-log"
)

func main() {
	log.DebugMode = true
	log.DebugSQLMode = true
	log.PrintTimestamp = true
	log.PrintColors = true
	log.TimeFormat = "2006-01-02 15:04:05.000"

	myVar := map[string]string{"hello": "world"}

	log.Debug("arg1", "arg2")
	log.Debugf("arg1 %d", 1)
	log.DebugDump(myVar, "prefix")
	log.DebugSeparator("title")
	log.DebugSQL("select * from mytable")

	log.Info("arg1", "arg2")
	log.Infof("arg1 %d", 1)
	log.InfoDump(myVar, "prefix")
	log.InfoSeparator("title")

	log.Warn("arg1", "arg2")
	log.Warnf("arg1 %d", 1)
	log.WarnDump(myVar, "prefix")
	log.WarnSeparator("title")

	log.Error("arg1", "arg2")
	log.Errorf("arg1 %d", 1)
	log.ErrorDump(myVar, "prefix")
	log.ErrorSeparator("title")

	log.Fatal("arg1", "arg2")
	log.Fatalf("arg1 %d", 1)

}
