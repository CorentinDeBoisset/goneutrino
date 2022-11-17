package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	buf := new(bytes.Buffer)
	SetLogOutput(buf)

	ErrorLog("some log with stuff: %s", "more stuff")
	logContent := buf.String()

	if len(logContent) != 60 {
		t.Fatalf("The generated log does not have the expected length (expected: %d, submitted: %d)", 60, len(logContent))
	}

	// Skip the date at the beginning of the log
	if logContent[19:] != " [ERROR] some log with stuff: more stuff\n" {
		t.Fatalf("The generated log has not the right content: %s", logContent[19:])
	}
}

func TestDefaultVerbosity(t *testing.T) {
	buf := new(bytes.Buffer)
	SetLogOutput(buf)

	ErrorLog("some error log")
	WarningLog("some warning log")
	InfoLog("some info log")
	DebugLog("some debug log")

	output := buf.String()
	lines := strings.Split(output, "\n")

	if len(lines) != 3 && !strings.Contains(lines[0], "some error log") || !strings.Contains(lines[1], "some warning log") {
		t.Fatalf("The generated log does not respect the verbosity")
	}
}

func TestVerbosityLevers(t *testing.T) {
	buf := new(bytes.Buffer)
	SetLogOutput(buf)

	SetLogLevel(1)
	if GetLogLevel() != 1 {
		t.Fatalf("The log level was not changed (expected: %d, actual: %d)", 1, GetLogLevel())
	}
	ErrorLog("some error log")
	WarningLog("some warning log")
	InfoLog("some info log")
	DebugLog("some debug log")

	output := buf.String()
	lines := strings.Split(output, "\n")

	if len(lines) != 4 && !strings.Contains(lines[2], "some info log") {
		t.Fatalf("The generated log does not respect the verbosity")
	}

	buf.Reset()
	SetLogLevel(2)
	if GetLogLevel() != 2 {
		t.Fatalf("The log level was not changed (expected: %d, actual: %d)", 2, GetLogLevel())
	}
	ErrorLog("some error log")
	WarningLog("some warning log")
	InfoLog("some info log")
	DebugLog("some debug log")

	output = buf.String()
	lines = strings.Split(output, "\n")

	if len(lines) != 5 && !strings.Contains(lines[3], "some debug log") {
		t.Fatalf("The generated log does not respect the verbosity")
	}
}
