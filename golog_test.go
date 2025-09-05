package golog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelTraceString(t *testing.T) {
	assert.Equal(t, "trace", LevelTrace.String())
}

func TestLevelDebugString(t *testing.T) {
	assert.Equal(t, "debug", LevelDebug.String())
}

func TestLevelInfoString(t *testing.T) {
	assert.Equal(t, "info", LevelInfo.String())
}

func TestLevelWarnString(t *testing.T) {
	assert.Equal(t, "warn", LevelWarn.String())
}

func TestLevelErrorString(t *testing.T) {
	assert.Equal(t, "error", LevelError.String())
}

func TestLevelFatalString(t *testing.T) {
	assert.Equal(t, "fatal", LevelFatal.String())
}

func TestLevelPanicString(t *testing.T) {
	assert.Equal(t, "panic", LevelPanic.String())
}

func TestLevelNoLevelString(t *testing.T) {
	assert.Equal(t, "", LevelNoLevel.String())
}

func TestParseLevelDebug(t *testing.T) {
	level, err := ParseLevel("DEBUG")
	assert.NoError(t, err)
	assert.Equal(t, LevelDebug, level)
}

func TestParseLevelInfo(t *testing.T) {
	level, err := ParseLevel("INFO")
	assert.NoError(t, err)
	assert.Equal(t, LevelInfo, level)
}

func TestParseLevelWarn(t *testing.T) {
	level, err := ParseLevel("WARN")
	assert.NoError(t, err)
	assert.Equal(t, LevelWarn, level)
}

func TestParseLevelError(t *testing.T) {
	level, err := ParseLevel("ERROR")
	assert.NoError(t, err)
	assert.Equal(t, LevelError, level)
}

func TestParseLevelFatal(t *testing.T) {
	level, err := ParseLevel("FATAL")
	assert.NoError(t, err)
	assert.Equal(t, LevelFatal, level)
}

func TestParseLevelPanic(t *testing.T) {
	level, err := ParseLevel("PANIC")
	assert.NoError(t, err)
	assert.Equal(t, LevelPanic, level)
}

func TestParseLevelLowerCase(t *testing.T) {
	level, err := ParseLevel("info")
	assert.NoError(t, err)
	assert.Equal(t, LevelInfo, level)
}
