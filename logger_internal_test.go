package log

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatMessage(t *testing.T) {

	input := []interface{}{"Hello %s\r\n%d 100%\r\n", "world", 1}
	expected := "Hello %s\r\n%d 100%\r\n world 1"
	actual := formatMessage(input...)

	assert.Equal(t, expected, actual)

}

func TestFormatSeparator(t *testing.T) {

	type test struct {
		name      string
		message   string
		separator string
		length    int
		expected  string
	}

	var tests = []test{
		{"no-message", "", "=", 12, "============"},
		{"short-length", "", "=", 2, "=="},
		{"short-length-with-message", "hello", "=", 2, "====[ hello ]"},
		{"short-message-1", "hello", "=", 12, "====[ hello ]"},
		{"short-message-2", "hello", "=", 20, "====[ hello ]======="},
		{"long-message-1", "hello world, how are you doing?", "=", 12, "====[ hello world, how are you doing? ]"},
		{"long-message-2", "hello world, how are you doing?", "=", 50, "====[ hello world, how are you doing? ]==========="},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := formatSeparator(tc.message, tc.separator, tc.length)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

func TestPrintMessage_NoColor(t *testing.T) {

	type test struct {
		name           string
		level          string
		message        string
		printTimestamp bool
		expectedStdout string
		expectedStderr string
	}

	var tests = []test{
		{"debug-1", "debug", "message", false, "message\n", ""},
		{"debug-2", "debug", "message", true, TestingTimeFormat + " | DEBUG | message\n", ""},

		{"info-1", "info ", "message", false, "message\n", ""},
		{"info-2", "info ", "message", true, TestingTimeFormat + " | INFO  | message\n", ""},

		{"warn-1", "warn ", "message", false, "message\n", ""},
		{"warn-2", "warn ", "message", true, TestingTimeFormat + " | WARN  | message\n", ""},

		{"error-1", "error", "message", false, "", "message\n"},
		{"error-2", "error", "message", true, "", TestingTimeFormat + " | ERROR | message\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			resetLogConfig()
			stdout, stderr := redirectOutput()
			defer resetLogOutput()

			PrintTimestamp = tc.printTimestamp
			PrintColors = false

			printMessage(tc.level, tc.message)

			actualStdOut := stdout.String()
			actualStdErr := stderr.String()

			assert.Equal(t, tc.expectedStdout, actualStdOut, "stdout")
			assert.Equal(t, tc.expectedStderr, actualStdErr, "stderr")

		})
	}

}

func TestPrintMessage_Color(t *testing.T) {

	type test struct {
		name           string
		level          string
		message        string
		printTimestamp bool
		expectedStdout string
		expectedStderr string
	}

	var tests = []test{
		{"debug-1", "debug", "message", false, "\x1b[90mmessage\x1b[0m\n", ""},
		{"debug-2", "debug", "message", true, "\x1b[90m" + TestingTimeFormat + " | DEBUG | message\x1b[0m\n", ""},

		{"info-1", "info ", "message", false, "\x1b[92mmessage\x1b[0m\n", ""},
		{"info-2", "info ", "message", true, "\x1b[92m" + TestingTimeFormat + " | INFO  | message\x1b[0m\n", ""},

		{"warn-1", "warn ", "message", false, "\x1b[93mmessage\x1b[0m\n", ""},
		{"warn-2", "warn ", "message", true, "\x1b[93m" + TestingTimeFormat + " | WARN  | message\x1b[0m\n", ""},

		{"error-1", "error", "message", false, "", "\x1b[91mmessage\x1b[0m\n"},
		{"error-2", "error", "message", true, "", "\x1b[91m" + TestingTimeFormat + " | ERROR | message\x1b[0m\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			resetLogConfig()
			stdout, stderr := redirectOutput()
			defer resetLogOutput()

			PrintTimestamp = tc.printTimestamp
			PrintColors = true

			printMessage(tc.level, tc.message)

			actualStdOut := stdout.String()
			actualStdErr := stderr.String()

			assert.Equal(t, tc.expectedStdout, actualStdOut, "stdout")
			assert.Equal(t, tc.expectedStderr, actualStdErr, "stderr")

		})
	}

}

func resetLogConfig() {
	PrintTimestamp = false
	DebugMode = false
	TimeZone, _ = time.LoadLocation("Europe/Brussels")
	TimeFormat = TestingTimeFormat
}

func redirectOutput() (*bytes.Buffer, *bytes.Buffer) {
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	Stdout = stdout
	Stderr = stderr
	return stdout, stderr
}

func resetLogOutput() {
	Stdout = os.Stdout
	Stderr = os.Stderr
	TimeFormat = DefaultTimeFormat
}
