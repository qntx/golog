package golog_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/qntx/golog"
)

// TestNew verifies that a new logger is created with default settings.
func TestNew(t *testing.T) {
	logger := golog.New()
	assert.NotNil(t, logger, "Logger should not be nil")
	assert.Equal(t, golog.LevelInfo, logger.GetLevel(), "Default level should be Info")
}

// TestNewWith verifies that a logger can be created with a custom zerolog.Logger.
func TestNewWith(t *testing.T) {
	var buf bytes.Buffer
	zerologLogger := zerolog.New(&buf).Level(zerolog.DebugLevel)
	l := golog.NewWith(zerologLogger)
	assert.NotNil(t, l, "Logger should not be nil")
	assert.Equal(t, golog.LevelDebug, l.GetLevel(), "Level should match custom logger")

	l.Debug("debug message")

	var logEntry map[string]any
	err := json.Unmarshal(buf.Bytes(), &logEntry)
	require.NoError(t, err, "Log output should be valid JSON")
	assert.Equal(t, "debug message", logEntry["message"], "Log message should match")
}

// TestLevel verifies dynamic level changes.
func TestLevel(t *testing.T) {
	l := golog.New()
	l.Level(golog.LevelDebug)
	assert.Equal(t, golog.LevelDebug, l.GetLevel(), "Level should be updated to Debug")

	var buf bytes.Buffer
	// Create a new golog.Logger that writes to the buffer for testing output filtering.
	// Initialize it with DebugLevel.
	bufferWritingZerolog := zerolog.New(&buf).Level(zerolog.DebugLevel)
	bufferLogger := golog.NewWith(bufferWritingZerolog)

	bufferLogger.Debug("debug message")
	assert.Contains(t, buf.String(), "debug message", "Debug message should be logged")

	// Change level to Warn and verify Debug messages are filtered.
	bufferLogger.Level(golog.LevelWarn)
	assert.Equal(t, golog.LevelWarn, bufferLogger.GetLevel(), "Buffer logger level should be updated to Warn")
	buf.Reset()
	bufferLogger.Debug("debug message")
	assert.Empty(t, buf.String(), "Debug message should not be logged at Warn level")
}

// TestParseLevel verifies level parsing from strings.
func TestParseLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected golog.Level
		hasError bool
	}{
		{"DEBUG", golog.LevelDebug, false},
		{"debug", golog.LevelDebug, false},
		{"INFO", golog.LevelInfo, false},
		{"WARN", golog.LevelWarn, false},
		{"ERROR", golog.LevelError, false},
		{"FATAL", golog.LevelFatal, false},
		{"PANIC", golog.LevelPanic, false},
		{"INVALID", golog.LevelInfo, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			level, err := golog.ParseLevel(tt.input)
			if tt.hasError {
				assert.Error(t, err, "Expected error for invalid level")
			} else {
				assert.NoError(t, err, "Expected no error for valid level")
				assert.Equal(t, tt.expected, level, "Parsed level should match")
			}
		})
	}
}

// TestLogLevels verifies logging at different levels.
func TestLogLevels(t *testing.T) {
	var buf bytes.Buffer
	logger := golog.NewWith(zerolog.New(&buf).Level(zerolog.TraceLevel))

	tests := []struct {
		logFunc func(string)
		level   string
		message string
	}{
		{logger.Trace, "trace", "trace message"},
		{logger.Debug, "debug", "debug message"},
		{logger.Info, "info", "info message"},
		{logger.Warn, "warn", "warn message"},
		{logger.Error, "error", "error message"},
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			buf.Reset()
			tt.logFunc(tt.message)

			var logEntry map[string]any
			err := json.Unmarshal(buf.Bytes(), &logEntry)
			require.NoError(t, err, "Log output should be valid JSON")
			assert.Equal(t, tt.level, logEntry["level"], "Log level should match")
			assert.Equal(t, tt.message, logEntry["message"], "Log message should match")
		})
	}
}

// TestFormattedLog verifies formatted logging methods.
func TestFormattedLog(t *testing.T) {
	var buf bytes.Buffer
	logger := golog.NewWith(zerolog.New(&buf).Level(zerolog.InfoLevel))

	logger.Infof("Hello, %s! You are %d years old.", "Alice", 30)

	var logEntry map[string]any
	err := json.Unmarshal(buf.Bytes(), &logEntry)
	require.NoError(t, err, "Log output should be valid JSON")
	assert.Equal(t, "Hello, Alice! You are 30 years old.", logEntry["message"], "Formatted message should match")
}

// TestPanic verifies that Panic logs and panics.
func TestPanic(t *testing.T) {
	var buf bytes.Buffer
	logger := golog.NewWith(zerolog.New(&buf).Level(zerolog.PanicLevel))

	// Capture panic.
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected panic but none occurred")
		}

		var logEntry map[string]any
		err := json.Unmarshal(buf.Bytes(), &logEntry)
		require.NoError(t, err, "Log output should be valid JSON")
		assert.Equal(t, "panic", logEntry["level"], "Log level should be panic")
		assert.Equal(t, "panic message", logEntry["message"], "Log message should match")
	}()

	logger.Panic("panic message")
}

// TestLevelString verifies the String method of Level.
func TestLevelString(t *testing.T) {
	tests := []struct {
		level    golog.Level
		expected string
	}{
		{golog.LevelTrace, "trace"},
		{golog.LevelDebug, "debug"},
		{golog.LevelInfo, "info"},
		{golog.LevelWarn, "warn"},
		{golog.LevelError, "error"},
		{golog.LevelFatal, "fatal"},
		{golog.LevelPanic, "panic"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.level.String(), "Level string should match")
		})
	}
}
