package log

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

var logMutex = &sync.Mutex{}

var colors = map[string]*color.Color{
	"DEBUG": color.New(color.FgHiBlack),
	"INFO ": color.New(color.FgHiGreen),
	"WARN ": color.New(color.FgHiYellow),
	"ERROR": color.New(color.FgHiRed),
	"FATAL": color.New(color.FgHiRed),
}

func init() {
	TimeZone, _ = time.LoadLocation("Europe/Brussels")
	DebugMode = os.Getenv("DEBUG") == "1"
	DebugSQLMode = os.Getenv("DEBUG_SQL") == "1"
	PrintTimestamp = os.Getenv("PRINT_TIMESTAMP") == "1"
}

func formatMessage(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	msg = strings.TrimRight(msg, " \n\r")
	return msg
}

func formatSeparator(message string, separator string, length int) string {
	if message == "" {
		return strings.Repeat(separator, length)
	}
	prefix := strings.Repeat(separator, 4)
	suffixLength := length - len(message) - len(prefix) - 4
	suffix := ""
	if suffixLength > 0 {
		suffix = strings.Repeat(separator, suffixLength)
	}
	return prefix + "[ " + message + " ]" + suffix
}

func printMessage(level string, message string) {

	logMutex.Lock()

	level = strings.ToUpper(level)

	if PrintTimestamp {
		tstamp := time.Now()
		if TimeZone != nil {
			tstamp = tstamp.In(TimeZone)
		}
		formattedTime := tstamp.Format(TimeFormat)
		message = formattedTime + " | " + level + " | " + message
	}
	if PrintColors {
		if color, ok := colors[level]; ok {
			message = color.Sprint(message)
		}
	}

	w := Stdout
	if level == "ERROR" || level == "FATAL" {
		w = Stderr
	}

	w.Write([]byte(message + "\n"))

	logMutex.Unlock()

}

func causeOfError(err error) error {

	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return err

}