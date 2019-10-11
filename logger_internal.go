package log

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var logMutex = &sync.Mutex{}

func init() {
	TimeZone, _ = time.LoadLocation("Europe/Brussels")
	DebugMode = os.Getenv("DEBUG") == "1"
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
		formattedTime := time.Now().In(TimeZone).Format(TimeFormat)
		message = formattedTime + " | " + level + " | " + message
	}

	w := Stdout
	if level == "ERROR" || level == "FATAL" {
		w = Stderr
	}

	fmt.Fprint(w, message+"\n")

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
