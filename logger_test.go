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

func Test_Debug_Enabled(t *testing.T) {

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

func Test_Debugf_Enabled(t *testing.T) {

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

func Test_Debug_Disabled(t *testing.T) {

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

func Test_DebugSQL_Enabled_Valid(t *testing.T) {

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

func Test_DebugSQL_Enabled_Error(t *testing.T) {

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

func Test_DebugSQL_Enabled_Empty(t *testing.T) {

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

func Test_DebugSQL_Disabled(t *testing.T) {

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

func Test_DebugSeparator_Disabled(t *testing.T) {

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

func Test_DebugSeparator_Enabled(t *testing.T) {

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

func Test_DebugDump_WithoutPrefix(t *testing.T) {

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

func Test_DebugDump_WithPrefix(t *testing.T) {

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

func Test_Info(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Info("info 100%")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | info 100%\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func Test_Infof(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Infof("info %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | info 2\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func Test_InfoSeparator(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.InfoSeparator("info")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | INFO  | ====[ info ]====================================================================\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func Test_InfoDump_WithoutPrefix(t *testing.T) {

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

func Test_InfoDump_WithPrefix(t *testing.T) {

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

func Test_Warn(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Warn("warn")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | warn\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func Test_Warnf(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Warnf("warn %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "test | WARN  | warn 2\n", actualStdOut, "stdout")
	assert.Equal(t, "", actualStdErr, "stderr")

}

func Test_WarnDump_WithoutPrefix(t *testing.T) {

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

func Test_WarnDump_WithPrefix(t *testing.T) {

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

func Test_Error(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Error("error")

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "test | ERROR | error\n", actualStdErr, "stderr")

}

func Test_Errorf(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.Errorf("error %d", 2)

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.Equal(t, "test | ERROR | error 2\n", actualStdErr, "stderr")

}

func Test_ErrorDump_WithoutPrefix(t *testing.T) {

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

func Test_ErrorDump_WithPrefix(t *testing.T) {

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

func Test_StackTrace(t *testing.T) {

	resetLogConfig()
	stdout, stderr := redirectOutput()
	defer resetLogOutput()

	log.StackTrace(errors.New("my error"))

	actualStdOut := stdout.String()
	actualStdErr := stderr.String()

	assert.Equal(t, "", actualStdOut, "stdout")
	assert.True(t, strings.HasPrefix(actualStdErr, "test | ERROR | *errors.fundamental my error\n"), "stderr")

}

func Test_FormattedStackTrace(t *testing.T) {
	actual := log.FormattedStackTrace(errors.New("my error"))
	assert.True(t, strings.HasPrefix(actual, "*errors.fundamental my error\n"))
}

func Test_Fatal(t *testing.T) {

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

func Test_Fatalf(t *testing.T) {

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

func Test_CheckError(t *testing.T) {

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
			if tc.debug == true {
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
