package log_test

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/pieterclaerhout/go-log"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestDebugEnabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.PrintColors = false

	log.Debug("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | debug\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugfEnabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.PrintColors = false

	log.Debugf("hello %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | hello 2\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugDisabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = false

	log.Debug("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugSQLEnabledValid(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.DebugSQLMode = true
	log.PrintColors = false

	log.DebugSQL("select * from mytable")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | SELECT *\nFROM mytable\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugSQLEnabledError(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.DebugSQLMode = true
	log.PrintColors = false

	log.DebugSQL("throw-error")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | Invalid SQL statement\n", actualStdErr)

}

func TestDebugSQLEnabledEmpty(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.DebugSQLMode = true
	log.PrintColors = false

	log.DebugSQL("")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | \n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugSQLDisabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = false
	log.DebugSQLMode = false

	log.DebugSQL("select * from mytable")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugSeparatorDisabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = false
	log.DebugSQLMode = false

	log.DebugSeparator("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugSeparatorEnabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.DebugSQLMode = true
	log.PrintColors = false

	log.DebugSeparator("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | ====[ debug ]===================================================================\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.PrintColors = false

	data := map[string]string{"hello": "world"}

	log.DebugDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestDebugDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.PrintColors = false

	data := map[string]string{"hello": "world"}

	log.DebugDump(data, "dprefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | dprefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestInfo(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.Info("info 100%")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | info 100%\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestInfof(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.Infof("info %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | info 2\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestInfoSeparator(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.InfoSeparator("info")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | ====[ info ]====================================================================\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestInfoDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	data := map[string]string{"hello": "world"}

	log.InfoDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestInfoDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	data := map[string]string{"hello": "world"}

	log.InfoDump(data, "iprefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | iprefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestWarn(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.Warn("warn")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | warn\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestWarnf(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.Warnf("warn %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | warn 2\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestWarnSeparator(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.WarnSeparator("info")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | ====[ info ]====================================================================\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestWarnDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	data := map[string]string{"hello": "world"}

	log.WarnDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestWarnDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	data := map[string]string{"hello": "world"}

	log.WarnDump(data, "wprefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | wprefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut)
	assert.Equal(t, "", actualStdErr)

}

func TestError(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.Error("error")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | error\n", actualStdErr)

}

func TestErrorf(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.Errorf("error %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | error 2\n", actualStdErr)

}

func TestErrorSeparator(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.ErrorSeparator("info")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | ====[ info ]====================================================================\n", actualStdErr)

}

func TestErrorDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	data := map[string]string{"hello": "world"}

	log.ErrorDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdErr)

}

func TestErrorDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	data := map[string]string{"hello": "world"}

	log.ErrorDump(data, "eprefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | ERROR | eprefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdErr)

}

func TestStackTrace(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.StackTrace(errors.New("my error"))

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.True(t, strings.HasPrefix(actualStdErr, "test | ERROR | my error\n"))
	assert.Equal(t, "test | ERROR | my error\n\tgo-log_test.TestStackTrace                        /Users/pclaerhout/Downloads/JonoFotografie/go-log/logger_test.go:521\n", actualStdErr)

}

type CustomError struct{}

func (m *CustomError) Error() string {
	return "boom"
}

func Test_StackTraceCustom(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	log.StackTrace(&CustomError{})

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "test | ERROR | boom\n\tgo-log_test.Test_StackTraceCustom                 /Users/pclaerhout/Downloads/JonoFotografie/go-log/logger_test.go:546\n", actualStdErr)

}

func TestFormattedStackTrace(t *testing.T) {
	actual := log.FormattedStackTrace(errors.New("my error"))
	assert.Equal(t, "my error\n\tgo-log_test.TestFormattedStackTrace               /Users/pclaerhout/Downloads/JonoFotografie/go-log/logger_test.go:557", actual)
}

func TestFatal(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	oldOsExit := log.OsExit
	defer func() {
		log.OsExit = oldOsExit
	}()

	var got int
	log.OsExit = func(code int) {
		got = code
	}

	log.Fatal("fatal error")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | FATAL | fatal error\n", actualStdErr)
	assert.Equal(t, 1, got)

}

func TestFatalf(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.PrintColors = false

	oldOsExit := log.OsExit
	defer func() {
		log.OsExit = oldOsExit
	}()

	var got int
	log.OsExit = func(code int) {
		got = code
	}

	log.Fatalf("fatal error %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut)
	assert.Equal(t, "test | FATAL | fatal error 2\n", actualStdErr)
	assert.Equal(t, 1, got)

}

func TestCheckError(t *testing.T) {

	type test struct {
		name             string
		err              error
		debug            bool
		expectedStdout   string
		expectedStderr   string
		expectedExitCode int
	}

	var tests = []test{
		{"nil-debug-nocolor", nil, true, "", "", -1},
		{"nil-debug-color", nil, true, "", "", -1},

		{"err-release-nocolor", errors.New("test"), false, "", "test | FATAL | test\n", 1},
		{"err-debug-nocolor", errors.New("test"), true, "", "test | FATAL | test\n", 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			resetLogConfig()
			stdout, stderr := redirectOutput()
			defer resetLogOutput()

			oldOsExit := log.OsExit
			defer func() {
				log.OsExit = oldOsExit
			}()

			var got int
			log.OsExit = func(code int) {
				got = code
			}

			log.DebugMode = tc.debug
			log.PrintColors = false

			log.CheckError(tc.err)

			actualStdOut := stdout.String()
			actualStdErr := stderr.String()

			assert.Equal(t, tc.expectedStdout, actualStdOut)
			if tc.debug {
				assert.True(t, strings.HasPrefix(actualStdErr, tc.expectedStderr), actualStdErr)
			} else {
				assert.Equal(t, tc.expectedStderr, actualStdErr)
			}

			if tc.expectedExitCode > 0 {
				assert.Equal(t, 1, got)
			}

		})
	}

}

func resetLogConfig() {
	log.PrintTimestamp = true
	log.PrintColors = true
	log.DebugMode = false
	log.DebugSQLMode = false
	log.TimeZone, _ = time.LoadLocation("Europe/Brussels")
	log.TimeFormat = log.TestingTimeFormat
}

func redirectOutput() (*bytes.Buffer, *bytes.Buffer) {
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	log.Stdout = stdout
	log.Stderr = stderr
	return stdout, stderr
}

func resetLogOutput() {
	log.Stdout = os.Stdout
	log.Stderr = os.Stderr
	log.TimeFormat = log.DefaultTimeFormat
}
