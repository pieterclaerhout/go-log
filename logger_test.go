package log_test

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-log"
)

func TestDebugEnabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true

	log.Debug("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | debug\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestDebugfEnabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true

	log.Debugf("hello %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | hello 2\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestDebugDisabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = false

	log.Debug("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestDebugSQLEnabledValid(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.DebugSQLMode = true

	log.DebugSQL("select * from mytable")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | SELECT *\nFROM mytable\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestDebugSQLEnabledError(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.DebugSQLMode = true

	log.DebugSQL("throw-error")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "test | ERROR | Invalid SQL statement\n", actualStdErr, "stderr")

}

func TestDebugSQLEnabledEmpty(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.DebugSQLMode = true

	log.DebugSQL("")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | \n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

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

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

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

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestDebugSeparatorEnabled(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true
	log.DebugSQLMode = true

	log.DebugSeparator("debug")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | ====[ debug ]===================================================================\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestDebugDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true

	data := map[string]string{"hello": "world"}

	log.DebugDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestDebugDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.DebugMode = true

	data := map[string]string{"hello": "world"}

	log.DebugDump(data, "prefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | DEBUG | prefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestInfo(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Info("info 100%")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | info 100%\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestInfof(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Infof("info %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | info 2\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestInfoSeparator(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.InfoSeparator("info")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | ====[ info ]====================================================================\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestInfoDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	data := map[string]string{"hello": "world"}

	log.InfoDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestInfoDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	data := map[string]string{"hello": "world"}

	log.InfoDump(data, "prefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | prefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestWarn(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Warn("warn")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | warn\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestWarnf(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Warnf("warn %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | warn 2\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestWarnDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	data := map[string]string{"hello": "world"}

	log.WarnDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestWarnDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	data := map[string]string{"hello": "world"}

	log.WarnDump(data, "prefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | prefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func TestError(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Error("error")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "test | ERROR | error\n", actualStdErr, "stderr")

}

func TestErrorf(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Errorf("error %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "test | ERROR | error 2\n", actualStdErr, "stderr")

}

func TestErrorDumpWithoutPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	data := map[string]string{"hello": "world"}

	log.ErrorDump(data, "")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "test | ERROR | map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdErr, "stderr")

}

func TestErrorDumpWithPrefix(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	data := map[string]string{"hello": "world"}

	log.ErrorDump(data, "prefix | ")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "test | ERROR | prefix |  map[string]string{\n  \"hello\": \"world\",\n}\n", actualStdErr, "stderr")

}

func TestStackTrace(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.StackTrace(errors.New("my error"))

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.True(t, strings.HasPrefix(actualStdErr, "test | ERROR | *errors.fundamental my error\n"), "stderr")

}

func TestFormattedStackTrace(t *testing.T) {
	actual := log.FormattedStackTrace(errors.New("my error"))
	assert.True(t, strings.HasPrefix(actual, "*errors.fundamental my error\n"))
}

func TestFatal(t *testing.T) {

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

	log.Fatal("fatal error")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.True(t, strings.HasPrefix(actualStdErr, "test | FATAL | fatal error\n"), "stderr")
	assert.Equal(t, 1, got, "exit-code")

}

func TestFatalf(t *testing.T) {

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

	log.Fatalf("fatal error %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.True(t, strings.HasPrefix(actualStdErr, "test | FATAL | fatal error 2\n"), "stderr")
	assert.Equal(t, 1, got, "exit-code")

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
		{"err-debug-nocolor", errors.New("test"), true, "", "test | FATAL | test\ntest | ERROR | *errors.fundamental test\n", 1},
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

			log.CheckError(tc.err)

			actualStdOut := stdout.String()
			actualStdErr := stderr.String()

			assert.Equal(t, tc.expectedStdout, actualStdOut, "stdout")
			if tc.debug {
				assert.True(t, strings.HasPrefix(actualStdErr, tc.expectedStderr), "stderr")
			} else {
				assert.Equal(t, tc.expectedStderr, actualStdErr, "stderr")
			}

			if tc.expectedExitCode > 0 {
				assert.Equal(t, 1, got, "exit-code")
			}

		})
	}

}

func resetLogConfig() {
	log.PrintTimestamp = true
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
